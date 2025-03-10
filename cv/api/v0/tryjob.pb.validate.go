// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: go.chromium.org/luci/cv/api/v0/tryjob.proto

package cvpb

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on Tryjob with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Tryjob) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Tryjob with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in TryjobMultiError, or nil if none found.
func (m *Tryjob) ValidateAll() error {
	return m.validate(true)
}

func (m *Tryjob) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Status

	if all {
		switch v := interface{}(m.GetResult()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TryjobValidationError{
					field:  "Result",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TryjobValidationError{
					field:  "Result",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetResult()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TryjobValidationError{
				field:  "Result",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Critical

	// no validation rules for Reuse

	if len(errors) > 0 {
		return TryjobMultiError(errors)
	}

	return nil
}

// TryjobMultiError is an error wrapping multiple validation errors returned by
// Tryjob.ValidateAll() if the designated constraints aren't met.
type TryjobMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TryjobMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TryjobMultiError) AllErrors() []error { return m }

// TryjobValidationError is the validation error returned by Tryjob.Validate if
// the designated constraints aren't met.
type TryjobValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TryjobValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TryjobValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TryjobValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TryjobValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TryjobValidationError) ErrorName() string { return "TryjobValidationError" }

// Error satisfies the builtin error interface
func (e TryjobValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTryjob.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TryjobValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TryjobValidationError{}

// Validate checks the field values on TryjobResult with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *TryjobResult) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TryjobResult with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in TryjobResultMultiError, or
// nil if none found.
func (m *TryjobResult) ValidateAll() error {
	return m.validate(true)
}

func (m *TryjobResult) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	switch v := m.Backend.(type) {
	case *TryjobResult_Buildbucket_:
		if v == nil {
			err := TryjobResultValidationError{
				field:  "Backend",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetBuildbucket()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, TryjobResultValidationError{
						field:  "Buildbucket",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, TryjobResultValidationError{
						field:  "Buildbucket",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetBuildbucket()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TryjobResultValidationError{
					field:  "Buildbucket",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return TryjobResultMultiError(errors)
	}

	return nil
}

// TryjobResultMultiError is an error wrapping multiple validation errors
// returned by TryjobResult.ValidateAll() if the designated constraints aren't met.
type TryjobResultMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TryjobResultMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TryjobResultMultiError) AllErrors() []error { return m }

// TryjobResultValidationError is the validation error returned by
// TryjobResult.Validate if the designated constraints aren't met.
type TryjobResultValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TryjobResultValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TryjobResultValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TryjobResultValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TryjobResultValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TryjobResultValidationError) ErrorName() string { return "TryjobResultValidationError" }

// Error satisfies the builtin error interface
func (e TryjobResultValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTryjobResult.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TryjobResultValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TryjobResultValidationError{}

// Validate checks the field values on TryjobInvocation with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *TryjobInvocation) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TryjobInvocation with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// TryjobInvocationMultiError, or nil if none found.
func (m *TryjobInvocation) ValidateAll() error {
	return m.validate(true)
}

func (m *TryjobInvocation) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetBuilderConfig()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TryjobInvocationValidationError{
					field:  "BuilderConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TryjobInvocationValidationError{
					field:  "BuilderConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetBuilderConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TryjobInvocationValidationError{
				field:  "BuilderConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Status

	// no validation rules for Critical

	for idx, item := range m.GetAttempts() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, TryjobInvocationValidationError{
						field:  fmt.Sprintf("Attempts[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, TryjobInvocationValidationError{
						field:  fmt.Sprintf("Attempts[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TryjobInvocationValidationError{
					field:  fmt.Sprintf("Attempts[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return TryjobInvocationMultiError(errors)
	}

	return nil
}

// TryjobInvocationMultiError is an error wrapping multiple validation errors
// returned by TryjobInvocation.ValidateAll() if the designated constraints
// aren't met.
type TryjobInvocationMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TryjobInvocationMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TryjobInvocationMultiError) AllErrors() []error { return m }

// TryjobInvocationValidationError is the validation error returned by
// TryjobInvocation.Validate if the designated constraints aren't met.
type TryjobInvocationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TryjobInvocationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TryjobInvocationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TryjobInvocationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TryjobInvocationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TryjobInvocationValidationError) ErrorName() string { return "TryjobInvocationValidationError" }

// Error satisfies the builtin error interface
func (e TryjobInvocationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTryjobInvocation.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TryjobInvocationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TryjobInvocationValidationError{}

// Validate checks the field values on Tryjob_Result with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Tryjob_Result) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Tryjob_Result with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in Tryjob_ResultMultiError, or
// nil if none found.
func (m *Tryjob_Result) ValidateAll() error {
	return m.validate(true)
}

func (m *Tryjob_Result) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Status

	switch v := m.Backend.(type) {
	case *Tryjob_Result_Buildbucket_:
		if v == nil {
			err := Tryjob_ResultValidationError{
				field:  "Backend",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetBuildbucket()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, Tryjob_ResultValidationError{
						field:  "Buildbucket",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, Tryjob_ResultValidationError{
						field:  "Buildbucket",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetBuildbucket()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return Tryjob_ResultValidationError{
					field:  "Buildbucket",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return Tryjob_ResultMultiError(errors)
	}

	return nil
}

// Tryjob_ResultMultiError is an error wrapping multiple validation errors
// returned by Tryjob_Result.ValidateAll() if the designated constraints
// aren't met.
type Tryjob_ResultMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Tryjob_ResultMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Tryjob_ResultMultiError) AllErrors() []error { return m }

// Tryjob_ResultValidationError is the validation error returned by
// Tryjob_Result.Validate if the designated constraints aren't met.
type Tryjob_ResultValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Tryjob_ResultValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Tryjob_ResultValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Tryjob_ResultValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Tryjob_ResultValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Tryjob_ResultValidationError) ErrorName() string { return "Tryjob_ResultValidationError" }

// Error satisfies the builtin error interface
func (e Tryjob_ResultValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTryjob_Result.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Tryjob_ResultValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Tryjob_ResultValidationError{}

// Validate checks the field values on Tryjob_Result_Buildbucket with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *Tryjob_Result_Buildbucket) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Tryjob_Result_Buildbucket with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// Tryjob_Result_BuildbucketMultiError, or nil if none found.
func (m *Tryjob_Result_Buildbucket) ValidateAll() error {
	return m.validate(true)
}

func (m *Tryjob_Result_Buildbucket) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return Tryjob_Result_BuildbucketMultiError(errors)
	}

	return nil
}

// Tryjob_Result_BuildbucketMultiError is an error wrapping multiple validation
// errors returned by Tryjob_Result_Buildbucket.ValidateAll() if the
// designated constraints aren't met.
type Tryjob_Result_BuildbucketMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Tryjob_Result_BuildbucketMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Tryjob_Result_BuildbucketMultiError) AllErrors() []error { return m }

// Tryjob_Result_BuildbucketValidationError is the validation error returned by
// Tryjob_Result_Buildbucket.Validate if the designated constraints aren't met.
type Tryjob_Result_BuildbucketValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Tryjob_Result_BuildbucketValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Tryjob_Result_BuildbucketValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Tryjob_Result_BuildbucketValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Tryjob_Result_BuildbucketValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Tryjob_Result_BuildbucketValidationError) ErrorName() string {
	return "Tryjob_Result_BuildbucketValidationError"
}

// Error satisfies the builtin error interface
func (e Tryjob_Result_BuildbucketValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTryjob_Result_Buildbucket.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Tryjob_Result_BuildbucketValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Tryjob_Result_BuildbucketValidationError{}

// Validate checks the field values on TryjobResult_Buildbucket with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *TryjobResult_Buildbucket) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TryjobResult_Buildbucket with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// TryjobResult_BuildbucketMultiError, or nil if none found.
func (m *TryjobResult_Buildbucket) ValidateAll() error {
	return m.validate(true)
}

func (m *TryjobResult_Buildbucket) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Host

	// no validation rules for Id

	if all {
		switch v := interface{}(m.GetBuilder()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TryjobResult_BuildbucketValidationError{
					field:  "Builder",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TryjobResult_BuildbucketValidationError{
					field:  "Builder",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetBuilder()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TryjobResult_BuildbucketValidationError{
				field:  "Builder",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return TryjobResult_BuildbucketMultiError(errors)
	}

	return nil
}

// TryjobResult_BuildbucketMultiError is an error wrapping multiple validation
// errors returned by TryjobResult_Buildbucket.ValidateAll() if the designated
// constraints aren't met.
type TryjobResult_BuildbucketMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TryjobResult_BuildbucketMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TryjobResult_BuildbucketMultiError) AllErrors() []error { return m }

// TryjobResult_BuildbucketValidationError is the validation error returned by
// TryjobResult_Buildbucket.Validate if the designated constraints aren't met.
type TryjobResult_BuildbucketValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TryjobResult_BuildbucketValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TryjobResult_BuildbucketValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TryjobResult_BuildbucketValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TryjobResult_BuildbucketValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TryjobResult_BuildbucketValidationError) ErrorName() string {
	return "TryjobResult_BuildbucketValidationError"
}

// Error satisfies the builtin error interface
func (e TryjobResult_BuildbucketValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTryjobResult_Buildbucket.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TryjobResult_BuildbucketValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TryjobResult_BuildbucketValidationError{}

// Validate checks the field values on TryjobInvocation_Attempt with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *TryjobInvocation_Attempt) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TryjobInvocation_Attempt with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// TryjobInvocation_AttemptMultiError, or nil if none found.
func (m *TryjobInvocation_Attempt) ValidateAll() error {
	return m.validate(true)
}

func (m *TryjobInvocation_Attempt) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Status

	if all {
		switch v := interface{}(m.GetResult()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TryjobInvocation_AttemptValidationError{
					field:  "Result",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TryjobInvocation_AttemptValidationError{
					field:  "Result",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetResult()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TryjobInvocation_AttemptValidationError{
				field:  "Result",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Reuse

	if len(errors) > 0 {
		return TryjobInvocation_AttemptMultiError(errors)
	}

	return nil
}

// TryjobInvocation_AttemptMultiError is an error wrapping multiple validation
// errors returned by TryjobInvocation_Attempt.ValidateAll() if the designated
// constraints aren't met.
type TryjobInvocation_AttemptMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TryjobInvocation_AttemptMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TryjobInvocation_AttemptMultiError) AllErrors() []error { return m }

// TryjobInvocation_AttemptValidationError is the validation error returned by
// TryjobInvocation_Attempt.Validate if the designated constraints aren't met.
type TryjobInvocation_AttemptValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TryjobInvocation_AttemptValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TryjobInvocation_AttemptValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TryjobInvocation_AttemptValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TryjobInvocation_AttemptValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TryjobInvocation_AttemptValidationError) ErrorName() string {
	return "TryjobInvocation_AttemptValidationError"
}

// Error satisfies the builtin error interface
func (e TryjobInvocation_AttemptValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTryjobInvocation_Attempt.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TryjobInvocation_AttemptValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TryjobInvocation_AttemptValidationError{}
