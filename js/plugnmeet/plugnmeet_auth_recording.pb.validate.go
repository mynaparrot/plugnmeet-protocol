// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: plugnmeet_auth_recording.proto

package plugnmeet

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

// Validate checks the field values on FetchRecordingsReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *FetchRecordingsReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FetchRecordingsReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// FetchRecordingsReqMultiError, or nil if none found.
func (m *FetchRecordingsReq) ValidateAll() error {
	return m.validate(true)
}

func (m *FetchRecordingsReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for From

	// no validation rules for Limit

	// no validation rules for OrderBy

	if len(errors) > 0 {
		return FetchRecordingsReqMultiError(errors)
	}

	return nil
}

// FetchRecordingsReqMultiError is an error wrapping multiple validation errors
// returned by FetchRecordingsReq.ValidateAll() if the designated constraints
// aren't met.
type FetchRecordingsReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FetchRecordingsReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FetchRecordingsReqMultiError) AllErrors() []error { return m }

// FetchRecordingsReqValidationError is the validation error returned by
// FetchRecordingsReq.Validate if the designated constraints aren't met.
type FetchRecordingsReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FetchRecordingsReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FetchRecordingsReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FetchRecordingsReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FetchRecordingsReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FetchRecordingsReqValidationError) ErrorName() string {
	return "FetchRecordingsReqValidationError"
}

// Error satisfies the builtin error interface
func (e FetchRecordingsReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFetchRecordingsReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FetchRecordingsReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FetchRecordingsReqValidationError{}

// Validate checks the field values on RecordingInfo with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *RecordingInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RecordingInfo with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in RecordingInfoMultiError, or
// nil if none found.
func (m *RecordingInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *RecordingInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for RecordId

	// no validation rules for RoomId

	// no validation rules for RoomSid

	// no validation rules for FilePath

	// no validation rules for FileSize

	// no validation rules for CreationTime

	// no validation rules for RoomCreationTime

	if len(errors) > 0 {
		return RecordingInfoMultiError(errors)
	}

	return nil
}

// RecordingInfoMultiError is an error wrapping multiple validation errors
// returned by RecordingInfo.ValidateAll() if the designated constraints
// aren't met.
type RecordingInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RecordingInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RecordingInfoMultiError) AllErrors() []error { return m }

// RecordingInfoValidationError is the validation error returned by
// RecordingInfo.Validate if the designated constraints aren't met.
type RecordingInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RecordingInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RecordingInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RecordingInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RecordingInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RecordingInfoValidationError) ErrorName() string { return "RecordingInfoValidationError" }

// Error satisfies the builtin error interface
func (e RecordingInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRecordingInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RecordingInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RecordingInfoValidationError{}

// Validate checks the field values on FetchRecordingsResult with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *FetchRecordingsResult) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FetchRecordingsResult with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// FetchRecordingsResultMultiError, or nil if none found.
func (m *FetchRecordingsResult) ValidateAll() error {
	return m.validate(true)
}

func (m *FetchRecordingsResult) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for TotalRecordings

	// no validation rules for From

	// no validation rules for Limit

	// no validation rules for OrderBy

	for idx, item := range m.GetRecordingsList() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, FetchRecordingsResultValidationError{
						field:  fmt.Sprintf("RecordingsList[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, FetchRecordingsResultValidationError{
						field:  fmt.Sprintf("RecordingsList[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return FetchRecordingsResultValidationError{
					field:  fmt.Sprintf("RecordingsList[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return FetchRecordingsResultMultiError(errors)
	}

	return nil
}

// FetchRecordingsResultMultiError is an error wrapping multiple validation
// errors returned by FetchRecordingsResult.ValidateAll() if the designated
// constraints aren't met.
type FetchRecordingsResultMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FetchRecordingsResultMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FetchRecordingsResultMultiError) AllErrors() []error { return m }

// FetchRecordingsResultValidationError is the validation error returned by
// FetchRecordingsResult.Validate if the designated constraints aren't met.
type FetchRecordingsResultValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FetchRecordingsResultValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FetchRecordingsResultValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FetchRecordingsResultValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FetchRecordingsResultValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FetchRecordingsResultValidationError) ErrorName() string {
	return "FetchRecordingsResultValidationError"
}

// Error satisfies the builtin error interface
func (e FetchRecordingsResultValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFetchRecordingsResult.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FetchRecordingsResultValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FetchRecordingsResultValidationError{}

// Validate checks the field values on FetchRecordingsRes with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *FetchRecordingsRes) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FetchRecordingsRes with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// FetchRecordingsResMultiError, or nil if none found.
func (m *FetchRecordingsRes) ValidateAll() error {
	return m.validate(true)
}

func (m *FetchRecordingsRes) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Status

	// no validation rules for Msg

	if all {
		switch v := interface{}(m.GetResult()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, FetchRecordingsResValidationError{
					field:  "Result",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, FetchRecordingsResValidationError{
					field:  "Result",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetResult()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FetchRecordingsResValidationError{
				field:  "Result",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return FetchRecordingsResMultiError(errors)
	}

	return nil
}

// FetchRecordingsResMultiError is an error wrapping multiple validation errors
// returned by FetchRecordingsRes.ValidateAll() if the designated constraints
// aren't met.
type FetchRecordingsResMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FetchRecordingsResMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FetchRecordingsResMultiError) AllErrors() []error { return m }

// FetchRecordingsResValidationError is the validation error returned by
// FetchRecordingsRes.Validate if the designated constraints aren't met.
type FetchRecordingsResValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FetchRecordingsResValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FetchRecordingsResValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FetchRecordingsResValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FetchRecordingsResValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FetchRecordingsResValidationError) ErrorName() string {
	return "FetchRecordingsResValidationError"
}

// Error satisfies the builtin error interface
func (e FetchRecordingsResValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFetchRecordingsRes.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FetchRecordingsResValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FetchRecordingsResValidationError{}

// Validate checks the field values on RecordingInfoReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *RecordingInfoReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RecordingInfoReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RecordingInfoReqMultiError, or nil if none found.
func (m *RecordingInfoReq) ValidateAll() error {
	return m.validate(true)
}

func (m *RecordingInfoReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for RecordId

	if len(errors) > 0 {
		return RecordingInfoReqMultiError(errors)
	}

	return nil
}

// RecordingInfoReqMultiError is an error wrapping multiple validation errors
// returned by RecordingInfoReq.ValidateAll() if the designated constraints
// aren't met.
type RecordingInfoReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RecordingInfoReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RecordingInfoReqMultiError) AllErrors() []error { return m }

// RecordingInfoReqValidationError is the validation error returned by
// RecordingInfoReq.Validate if the designated constraints aren't met.
type RecordingInfoReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RecordingInfoReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RecordingInfoReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RecordingInfoReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RecordingInfoReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RecordingInfoReqValidationError) ErrorName() string { return "RecordingInfoReqValidationError" }

// Error satisfies the builtin error interface
func (e RecordingInfoReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRecordingInfoReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RecordingInfoReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RecordingInfoReqValidationError{}

// Validate checks the field values on RecordingInfoRes with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *RecordingInfoRes) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RecordingInfoRes with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RecordingInfoResMultiError, or nil if none found.
func (m *RecordingInfoRes) ValidateAll() error {
	return m.validate(true)
}

func (m *RecordingInfoRes) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Status

	// no validation rules for Msg

	if all {
		switch v := interface{}(m.GetRecordingInfo()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RecordingInfoResValidationError{
					field:  "RecordingInfo",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RecordingInfoResValidationError{
					field:  "RecordingInfo",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRecordingInfo()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RecordingInfoResValidationError{
				field:  "RecordingInfo",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetRoomInfo()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RecordingInfoResValidationError{
					field:  "RoomInfo",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RecordingInfoResValidationError{
					field:  "RoomInfo",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRoomInfo()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RecordingInfoResValidationError{
				field:  "RoomInfo",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return RecordingInfoResMultiError(errors)
	}

	return nil
}

// RecordingInfoResMultiError is an error wrapping multiple validation errors
// returned by RecordingInfoRes.ValidateAll() if the designated constraints
// aren't met.
type RecordingInfoResMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RecordingInfoResMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RecordingInfoResMultiError) AllErrors() []error { return m }

// RecordingInfoResValidationError is the validation error returned by
// RecordingInfoRes.Validate if the designated constraints aren't met.
type RecordingInfoResValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RecordingInfoResValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RecordingInfoResValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RecordingInfoResValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RecordingInfoResValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RecordingInfoResValidationError) ErrorName() string { return "RecordingInfoResValidationError" }

// Error satisfies the builtin error interface
func (e RecordingInfoResValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRecordingInfoRes.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RecordingInfoResValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RecordingInfoResValidationError{}

// Validate checks the field values on DeleteRecordingReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteRecordingReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteRecordingReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteRecordingReqMultiError, or nil if none found.
func (m *DeleteRecordingReq) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteRecordingReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for RecordId

	if len(errors) > 0 {
		return DeleteRecordingReqMultiError(errors)
	}

	return nil
}

// DeleteRecordingReqMultiError is an error wrapping multiple validation errors
// returned by DeleteRecordingReq.ValidateAll() if the designated constraints
// aren't met.
type DeleteRecordingReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteRecordingReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteRecordingReqMultiError) AllErrors() []error { return m }

// DeleteRecordingReqValidationError is the validation error returned by
// DeleteRecordingReq.Validate if the designated constraints aren't met.
type DeleteRecordingReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteRecordingReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteRecordingReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteRecordingReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteRecordingReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteRecordingReqValidationError) ErrorName() string {
	return "DeleteRecordingReqValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteRecordingReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteRecordingReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteRecordingReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteRecordingReqValidationError{}

// Validate checks the field values on DeleteRecordingRes with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteRecordingRes) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteRecordingRes with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteRecordingResMultiError, or nil if none found.
func (m *DeleteRecordingRes) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteRecordingRes) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Status

	// no validation rules for Msg

	if len(errors) > 0 {
		return DeleteRecordingResMultiError(errors)
	}

	return nil
}

// DeleteRecordingResMultiError is an error wrapping multiple validation errors
// returned by DeleteRecordingRes.ValidateAll() if the designated constraints
// aren't met.
type DeleteRecordingResMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteRecordingResMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteRecordingResMultiError) AllErrors() []error { return m }

// DeleteRecordingResValidationError is the validation error returned by
// DeleteRecordingRes.Validate if the designated constraints aren't met.
type DeleteRecordingResValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteRecordingResValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteRecordingResValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteRecordingResValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteRecordingResValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteRecordingResValidationError) ErrorName() string {
	return "DeleteRecordingResValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteRecordingResValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteRecordingRes.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteRecordingResValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteRecordingResValidationError{}

// Validate checks the field values on GetDownloadTokenReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetDownloadTokenReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetDownloadTokenReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetDownloadTokenReqMultiError, or nil if none found.
func (m *GetDownloadTokenReq) ValidateAll() error {
	return m.validate(true)
}

func (m *GetDownloadTokenReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for RecordId

	if len(errors) > 0 {
		return GetDownloadTokenReqMultiError(errors)
	}

	return nil
}

// GetDownloadTokenReqMultiError is an error wrapping multiple validation
// errors returned by GetDownloadTokenReq.ValidateAll() if the designated
// constraints aren't met.
type GetDownloadTokenReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetDownloadTokenReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetDownloadTokenReqMultiError) AllErrors() []error { return m }

// GetDownloadTokenReqValidationError is the validation error returned by
// GetDownloadTokenReq.Validate if the designated constraints aren't met.
type GetDownloadTokenReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetDownloadTokenReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetDownloadTokenReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetDownloadTokenReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetDownloadTokenReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetDownloadTokenReqValidationError) ErrorName() string {
	return "GetDownloadTokenReqValidationError"
}

// Error satisfies the builtin error interface
func (e GetDownloadTokenReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetDownloadTokenReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetDownloadTokenReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetDownloadTokenReqValidationError{}

// Validate checks the field values on GetDownloadTokenRes with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetDownloadTokenRes) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetDownloadTokenRes with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetDownloadTokenResMultiError, or nil if none found.
func (m *GetDownloadTokenRes) ValidateAll() error {
	return m.validate(true)
}

func (m *GetDownloadTokenRes) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Status

	// no validation rules for Msg

	if m.Token != nil {
		// no validation rules for Token
	}

	if len(errors) > 0 {
		return GetDownloadTokenResMultiError(errors)
	}

	return nil
}

// GetDownloadTokenResMultiError is an error wrapping multiple validation
// errors returned by GetDownloadTokenRes.ValidateAll() if the designated
// constraints aren't met.
type GetDownloadTokenResMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetDownloadTokenResMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetDownloadTokenResMultiError) AllErrors() []error { return m }

// GetDownloadTokenResValidationError is the validation error returned by
// GetDownloadTokenRes.Validate if the designated constraints aren't met.
type GetDownloadTokenResValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetDownloadTokenResValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetDownloadTokenResValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetDownloadTokenResValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetDownloadTokenResValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetDownloadTokenResValidationError) ErrorName() string {
	return "GetDownloadTokenResValidationError"
}

// Error satisfies the builtin error interface
func (e GetDownloadTokenResValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetDownloadTokenRes.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetDownloadTokenResValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetDownloadTokenResValidationError{}
