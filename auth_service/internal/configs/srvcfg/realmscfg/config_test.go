// Copyright 2023 The LUCI Authors.
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

package realmscfg

import (
	"context"
	"testing"

	"go.chromium.org/luci/common/proto/realms"
	"go.chromium.org/luci/common/testing/ftt"
	"go.chromium.org/luci/common/testing/truth/assert"
	"go.chromium.org/luci/common/testing/truth/should"
	"go.chromium.org/luci/gae/impl/memory"
)

func TestConfigContext(t *testing.T) {
	t.Parallel()
	ctx := memory.Use(context.Background())
	realmsCfg := &realms.RealmsCfg{
		Realms: []*realms.Realm{
			{
				Name: "test-realm",
				Extends: []string{
					"another-test-realm",
				},
				Bindings: []*realms.Binding{
					{
						Role: "role/test.serviceAccountTokenCreator",
						Principals: []string{
							"group:test-reader-only",
							"user:test-service-account-uploader@example.service.com",
						},
					},
				},
			},
		},
	}

	ftt.Run("Getting without setting fails", t, func(t *ftt.Test) {
		_, err := Get(ctx)
		assert.Loosely(t, err, should.NotBeNil)
	})

	ftt.Run("Testing basic config operations", t, func(t *ftt.Test) {
		assert.Loosely(t, SetConfig(ctx, realmsCfg), should.BeNil)
		cfgFromGet, err := Get(ctx)
		assert.Loosely(t, err, should.BeNil)
		assert.Loosely(t, cfgFromGet, should.Match(realmsCfg))
	})
}
