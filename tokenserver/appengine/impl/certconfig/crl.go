// Copyright 2016 The LUCI Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package certconfig

import (
	"context"
	"crypto/sha1"
	"crypto/x509/pkix"
	"encoding/hex"
	"fmt"
	"math/big"
	"sync"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	"go.chromium.org/luci/common/data/caching/lazyslot"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/common/retry/transient"
	ds "go.chromium.org/luci/gae/service/datastore"

	"go.chromium.org/luci/tokenserver/api/admin/v1"
	"go.chromium.org/luci/tokenserver/appengine/impl/utils"
	"go.chromium.org/luci/tokenserver/appengine/impl/utils/shards"
)

// CRLShardCount is a number of shards to use for storing CRL in the datastore.
//
// Each shard can hold ~2 MB of data (taking into account zlib compression),
// so 16 shards ~= 32 MB. Good enough for a foreseeable future.
//
// Changing this value requires rerunning of Admin.FetchCRL RPC to rebuild
// the entities.
const CRLShardCount = 16

// CRL represents a parsed Certificate Revocation List of some CA.
//
// ID is always "crl", the parent entity is corresponding CA.
type CRL struct {
	_id string `gae:"$id,crl"`

	// Parent is pointing to parent CA entity.
	Parent *ds.Key `gae:"$parent"`

	// EntityVersion is used for simple concurrency control.
	//
	// Increase on each update of this entity.
	EntityVersion int `gae:",noindex"`

	// LastUpdateTime is extracted from corresponding field of CRL.
	//
	// It indicates a time when CRL was generated by the CA.
	LastUpdateTime time.Time `gae:",noindex"`

	// LastFetchTime is when this CRL was fetched the last time.
	//
	// Updated only when newer CRL version is fetched.
	LastFetchTime time.Time `gae:",noindex"`

	// LastFetchETag is ETag header of last downloaded CRL file.
	//
	// If CRL's etag doesn't change, we can skip reparsing CRL.
	LastFetchETag string `gae:",noindex"`

	// RevokedCertsCount is a number of revoked certificates in CRL. FYI only.
	RevokedCertsCount int `gae:",noindex"`
}

// GetStatusProto returns populated CRLStatus proto message.
func (crl *CRL) GetStatusProto() *admin.CRLStatus {
	return &admin.CRLStatus{
		LastUpdateTime:    timestamppb.New(crl.LastUpdateTime),
		LastFetchTime:     timestamppb.New(crl.LastFetchTime),
		LastFetchEtag:     crl.LastFetchETag,
		RevokedCertsCount: int64(crl.RevokedCertsCount),
	}
}

////////////////////////////////////////////////////////////////////////////////

// CRLShardHeader represents a hash of a shard of a CRL sharded set.
//
// We split CRL into a bunch of shards to avoid hitting datastore entity size
// limits. Each shard lives in its own entity group, where root entity
// (CRLShardHeader) contains a hash of the shard data (CRLShardBody).
//
// It is used to skip fetches of fat shard entities if we already have the same
// data locally (based on matching hash).
//
// ID is "<cn name>|<total number of shards>|<shard index>" (see shardEntityID).
type CRLShardHeader struct {
	ID   string `gae:"$id"`
	SHA1 string `gae:",noindex"` // SHA1 of serialized shard data (before compression)
}

// CRLShardBody is a fat entity that contains serialized CRL shard.
//
// See CRLShardHeader for more info.
//
// Parent entity is CRLShardHeader. ID is always "1".
type CRLShardBody struct {
	_id string `gae:"$id,1"`

	Parent     *ds.Key `gae:"$parent"`  // key of CRLShardHeader
	SHA1       string  `gae:",noindex"` // SHA1 of serialized shard data (before compression)
	ZippedData []byte  `gae:",noindex"` // zlib-compressed serialized shards.Shard.
}

// UpdateCRLSet splits a set of revoked certificate serial numbers into shards,
// storing each shard in a separate entity (CRLShardBody).
//
// It effectively overwrites the entire set.
func UpdateCRLSet(c context.Context, cn string, shardCount int, crl *pkix.CertificateList) error {
	// Split CRL into shards.
	set := make(shards.Set, shardCount)
	for _, cert := range crl.TBSCertList.RevokedCertificates {
		sn, err := utils.SerializeSN(cert.SerialNumber)
		if err != nil {
			return err
		}
		set.Insert(sn)
	}
	// Update shards in parallel via a bunch of independent transactions.
	wg := sync.WaitGroup{}
	er := errors.NewLazyMultiError(len(set))
	for idx, shard := range set {
		wg.Add(1)
		go func(idx int, shard shards.Shard) {
			defer wg.Done()
			er.Assign(idx, updateCRLShard(c, cn, shard, shardCount, idx))
		}(idx, shard)
	}
	wg.Wait()
	return er.Get()
}

// updateCRLShard updates entities that holds a single shard of a CRL set.
func updateCRLShard(c context.Context, cn string, shard shards.Shard, count, idx int) error {
	blob := shard.Serialize()
	hash := sha1.Sum(blob)
	digest := hex.EncodeToString(hash[:])

	// Have it already?
	header := CRLShardHeader{ID: shardEntityID(cn, count, idx)}
	switch err := ds.Get(c, &header); {
	case err != nil && err != ds.ErrNoSuchEntity:
		return err
	case err == nil && header.SHA1 == digest:
		logging.Infof(c, "CRL for %q: shard %d/%d is up-to-date", cn, idx, count)
		return nil
	}

	// Zip before uploading.
	zipped, err := utils.ZlibCompress(blob)
	if err != nil {
		return err
	}
	logging.Infof(
		c, "CRL for %q: shard %d/%d updated (%d bytes zipped, %d%% compression)",
		cn, idx, count, len(zipped), 100*len(zipped)/len(blob))

	// Upload, updating the header and the body at once.
	return ds.RunInTransaction(c, func(c context.Context) error {
		header.SHA1 = digest
		body := CRLShardBody{
			Parent:     ds.KeyForObj(c, &header),
			SHA1:       digest,
			ZippedData: zipped,
		}
		return ds.Put(c, &header, &body)
	}, nil)
}

// shardEntityID returns an ID of CRLShardHeader entity for given shard.
//
// 'cn' is Common Name of the CRL. 'total' is total number of shards expected,
// and 'index' is an index of some particular shard.
func shardEntityID(cn string, total, index int) string {
	return fmt.Sprintf("%s|%d|%d", cn, total, index)
}

////////////////////////////////////////////////////////////////////////////////

// CRLChecker knows how to check presence of a certificate serial number in CRL.
//
// Uses entities prepared by UpdateCRLSet.
//
// It is a stateful object that caches CRL shards in memory (occasionally
// refetching them from the datastore), thus providing an eventually consistent
// view of the CRL set.
//
// Safe for concurrent use. Should be reused between requests.
type CRLChecker struct {
	cn            string          // name of CA to check a CRL of
	shardCount    int             // a total number of shards
	shards        []lazyslot.Slot // per-shard local state, len(shards) == shardCount
	cacheDuration time.Duration   // how often to refetch shards from datastore
}

// shardCache is kept inside 'shards' slots in CRLChecker.
type shardCache struct {
	shard shards.Shard // shard data as a map[]
	sha1  string       // shard hash, to skip unnecessary refetches
}

// NewCRLChecker initializes new CRLChecker that knows how to examine CRL of
// a CA (identifies by its Common Name).
//
// It must know number of shards in advance. Usually is it just CRLShardCount.
//
// It will cache shards in local memory, refetching them if necessary after
// 'cacheDuration' interval.
func NewCRLChecker(cn string, shardCount int, cacheDuration time.Duration) *CRLChecker {
	return &CRLChecker{
		cn:            cn,
		shardCount:    shardCount,
		shards:        make([]lazyslot.Slot, shardCount),
		cacheDuration: cacheDuration,
	}
}

// IsRevokedSN returns true if given serial number is in the CRL.
func (ch *CRLChecker) IsRevokedSN(c context.Context, sn *big.Int) (bool, error) {
	snBlob, err := utils.SerializeSN(sn)
	if err != nil {
		return false, err
	}
	shard, err := ch.shard(c, shards.ShardIndex(snBlob, ch.shardCount))
	if err != nil {
		return false, err
	}
	_, revoked := shard[string(snBlob)]
	return revoked, nil
}

// shard returns a shard given its index.
func (ch *CRLChecker) shard(c context.Context, idx int) (shards.Shard, error) {
	val, err := ch.shards[idx].Get(c, func(prev any) (any, time.Duration, error) {
		prevState, _ := prev.(shardCache)
		newState, err := ch.refetchShard(c, idx, prevState)
		return newState, ch.cacheDuration, err
	})
	if err != nil {
		return nil, err
	}
	// lazyslot.Get always returns non-nil val on success. It is safe to cast it
	// to whatever we returned in the callback (which is always shardCache, see
	// refetchShard).
	return val.(shardCache).shard, nil
}

// refetchShard is called by 'shard' to fetch a new version of a shard.
func (ch *CRLChecker) refetchShard(c context.Context, idx int, prevState shardCache) (newState shardCache, err error) {
	// Have something locally already? Quickly fetch CRLShardHeader to check
	// whether we need to pull a heavy CRLShardBody.
	hdr := CRLShardHeader{ID: shardEntityID(ch.cn, ch.shardCount, idx)}
	if prevState.sha1 != "" {
		switch err = ds.Get(c, &hdr); {
		case err == ds.ErrNoSuchEntity:
			err = fmt.Errorf("shard header %q is missing", hdr.ID)
			return
		case err != nil:
			err = transient.Tag.Apply(err)
			return
		}
		// The currently cached copy is still good enough?
		if hdr.SHA1 == prevState.sha1 {
			newState = prevState
			return
		}
	}

	// Nothing is cached, or the datastore copy is fresher than what we have in
	// the cache. Need to fetch a new copy, unzip and deserialize it. This entity
	// is prepared by updateCRLShard.
	body := CRLShardBody{Parent: ds.KeyForObj(c, &hdr)}
	switch err = ds.Get(c, &body); {
	case err == ds.ErrNoSuchEntity:
		err = fmt.Errorf("shard body %q is missing", hdr.ID)
		return
	case err != nil:
		err = transient.Tag.Apply(err)
		return
	}

	// Unzip and deserialize.
	blob, err := utils.ZlibDecompress(body.ZippedData)
	if err != nil {
		return
	}
	shard, err := shards.ParseShard(blob)
	if err != nil {
		return
	}

	newState = shardCache{shard: shard, sha1: body.SHA1}
	return
}
