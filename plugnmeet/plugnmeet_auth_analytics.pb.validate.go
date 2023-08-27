// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: plugnmeet_auth_analytics.proto

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

// Validate checks the field values on FetchAnalyticsReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *FetchAnalyticsReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FetchAnalyticsReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// FetchAnalyticsReqMultiError, or nil if none found.
func (m *FetchAnalyticsReq) ValidateAll() error {
	return m.validate(true)
}

func (m *FetchAnalyticsReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for From

	// no validation rules for Limit

	// no validation rules for OrderBy

	if len(errors) > 0 {
		return FetchAnalyticsReqMultiError(errors)
	}

	return nil
}

// FetchAnalyticsReqMultiError is an error wrapping multiple validation errors
// returned by FetchAnalyticsReq.ValidateAll() if the designated constraints
// aren't met.
type FetchAnalyticsReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FetchAnalyticsReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FetchAnalyticsReqMultiError) AllErrors() []error { return m }

// FetchAnalyticsReqValidationError is the validation error returned by
// FetchAnalyticsReq.Validate if the designated constraints aren't met.
type FetchAnalyticsReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FetchAnalyticsReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FetchAnalyticsReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FetchAnalyticsReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FetchAnalyticsReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FetchAnalyticsReqValidationError) ErrorName() string {
	return "FetchAnalyticsReqValidationError"
}

// Error satisfies the builtin error interface
func (e FetchAnalyticsReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFetchAnalyticsReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FetchAnalyticsReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FetchAnalyticsReqValidationError{}

// Validate checks the field values on AnalyticsInfo with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AnalyticsInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AnalyticsInfo with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AnalyticsInfoMultiError, or
// nil if none found.
func (m *AnalyticsInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *AnalyticsInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for RoomId

	// no validation rules for FileId

	// no validation rules for FileName

	// no validation rules for CreationTime

	// no validation rules for RoomCreationTime

	if len(errors) > 0 {
		return AnalyticsInfoMultiError(errors)
	}

	return nil
}

// AnalyticsInfoMultiError is an error wrapping multiple validation errors
// returned by AnalyticsInfo.ValidateAll() if the designated constraints
// aren't met.
type AnalyticsInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AnalyticsInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AnalyticsInfoMultiError) AllErrors() []error { return m }

// AnalyticsInfoValidationError is the validation error returned by
// AnalyticsInfo.Validate if the designated constraints aren't met.
type AnalyticsInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AnalyticsInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AnalyticsInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AnalyticsInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AnalyticsInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AnalyticsInfoValidationError) ErrorName() string { return "AnalyticsInfoValidationError" }

// Error satisfies the builtin error interface
func (e AnalyticsInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAnalyticsInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AnalyticsInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AnalyticsInfoValidationError{}

// Validate checks the field values on FetchAnalyticsResult with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *FetchAnalyticsResult) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FetchAnalyticsResult with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// FetchAnalyticsResultMultiError, or nil if none found.
func (m *FetchAnalyticsResult) ValidateAll() error {
	return m.validate(true)
}

func (m *FetchAnalyticsResult) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for TotalAnalytics

	// no validation rules for From

	// no validation rules for Limit

	// no validation rules for OrderBy

	for idx, item := range m.GetAnalyticsList() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, FetchAnalyticsResultValidationError{
						field:  fmt.Sprintf("AnalyticsList[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, FetchAnalyticsResultValidationError{
						field:  fmt.Sprintf("AnalyticsList[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return FetchAnalyticsResultValidationError{
					field:  fmt.Sprintf("AnalyticsList[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return FetchAnalyticsResultMultiError(errors)
	}

	return nil
}

// FetchAnalyticsResultMultiError is an error wrapping multiple validation
// errors returned by FetchAnalyticsResult.ValidateAll() if the designated
// constraints aren't met.
type FetchAnalyticsResultMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FetchAnalyticsResultMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FetchAnalyticsResultMultiError) AllErrors() []error { return m }

// FetchAnalyticsResultValidationError is the validation error returned by
// FetchAnalyticsResult.Validate if the designated constraints aren't met.
type FetchAnalyticsResultValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FetchAnalyticsResultValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FetchAnalyticsResultValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FetchAnalyticsResultValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FetchAnalyticsResultValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FetchAnalyticsResultValidationError) ErrorName() string {
	return "FetchAnalyticsResultValidationError"
}

// Error satisfies the builtin error interface
func (e FetchAnalyticsResultValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFetchAnalyticsResult.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FetchAnalyticsResultValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FetchAnalyticsResultValidationError{}

// Validate checks the field values on FetchAnalyticsRes with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *FetchAnalyticsRes) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FetchAnalyticsRes with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// FetchAnalyticsResMultiError, or nil if none found.
func (m *FetchAnalyticsRes) ValidateAll() error {
	return m.validate(true)
}

func (m *FetchAnalyticsRes) validate(all bool) error {
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
				errors = append(errors, FetchAnalyticsResValidationError{
					field:  "Result",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, FetchAnalyticsResValidationError{
					field:  "Result",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetResult()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FetchAnalyticsResValidationError{
				field:  "Result",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return FetchAnalyticsResMultiError(errors)
	}

	return nil
}

// FetchAnalyticsResMultiError is an error wrapping multiple validation errors
// returned by FetchAnalyticsRes.ValidateAll() if the designated constraints
// aren't met.
type FetchAnalyticsResMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FetchAnalyticsResMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FetchAnalyticsResMultiError) AllErrors() []error { return m }

// FetchAnalyticsResValidationError is the validation error returned by
// FetchAnalyticsRes.Validate if the designated constraints aren't met.
type FetchAnalyticsResValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FetchAnalyticsResValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FetchAnalyticsResValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FetchAnalyticsResValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FetchAnalyticsResValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FetchAnalyticsResValidationError) ErrorName() string {
	return "FetchAnalyticsResValidationError"
}

// Error satisfies the builtin error interface
func (e FetchAnalyticsResValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFetchAnalyticsRes.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FetchAnalyticsResValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FetchAnalyticsResValidationError{}

// Validate checks the field values on DeleteAnalyticsReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteAnalyticsReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteAnalyticsReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteAnalyticsReqMultiError, or nil if none found.
func (m *DeleteAnalyticsReq) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteAnalyticsReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetFileId()) < 1 {
		err := DeleteAnalyticsReqValidationError{
			field:  "FileId",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return DeleteAnalyticsReqMultiError(errors)
	}

	return nil
}

// DeleteAnalyticsReqMultiError is an error wrapping multiple validation errors
// returned by DeleteAnalyticsReq.ValidateAll() if the designated constraints
// aren't met.
type DeleteAnalyticsReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteAnalyticsReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteAnalyticsReqMultiError) AllErrors() []error { return m }

// DeleteAnalyticsReqValidationError is the validation error returned by
// DeleteAnalyticsReq.Validate if the designated constraints aren't met.
type DeleteAnalyticsReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteAnalyticsReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteAnalyticsReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteAnalyticsReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteAnalyticsReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteAnalyticsReqValidationError) ErrorName() string {
	return "DeleteAnalyticsReqValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteAnalyticsReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteAnalyticsReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteAnalyticsReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteAnalyticsReqValidationError{}

// Validate checks the field values on DeleteAnalyticsRes with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteAnalyticsRes) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteAnalyticsRes with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteAnalyticsResMultiError, or nil if none found.
func (m *DeleteAnalyticsRes) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteAnalyticsRes) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Status

	// no validation rules for Msg

	if len(errors) > 0 {
		return DeleteAnalyticsResMultiError(errors)
	}

	return nil
}

// DeleteAnalyticsResMultiError is an error wrapping multiple validation errors
// returned by DeleteAnalyticsRes.ValidateAll() if the designated constraints
// aren't met.
type DeleteAnalyticsResMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteAnalyticsResMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteAnalyticsResMultiError) AllErrors() []error { return m }

// DeleteAnalyticsResValidationError is the validation error returned by
// DeleteAnalyticsRes.Validate if the designated constraints aren't met.
type DeleteAnalyticsResValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteAnalyticsResValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteAnalyticsResValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteAnalyticsResValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteAnalyticsResValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteAnalyticsResValidationError) ErrorName() string {
	return "DeleteAnalyticsResValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteAnalyticsResValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteAnalyticsRes.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteAnalyticsResValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteAnalyticsResValidationError{}

// Validate checks the field values on GetAnalyticsDownloadTokenReq with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetAnalyticsDownloadTokenReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetAnalyticsDownloadTokenReq with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetAnalyticsDownloadTokenReqMultiError, or nil if none found.
func (m *GetAnalyticsDownloadTokenReq) ValidateAll() error {
	return m.validate(true)
}

func (m *GetAnalyticsDownloadTokenReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetFileId()) < 1 {
		err := GetAnalyticsDownloadTokenReqValidationError{
			field:  "FileId",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetAnalyticsDownloadTokenReqMultiError(errors)
	}

	return nil
}

// GetAnalyticsDownloadTokenReqMultiError is an error wrapping multiple
// validation errors returned by GetAnalyticsDownloadTokenReq.ValidateAll() if
// the designated constraints aren't met.
type GetAnalyticsDownloadTokenReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetAnalyticsDownloadTokenReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetAnalyticsDownloadTokenReqMultiError) AllErrors() []error { return m }

// GetAnalyticsDownloadTokenReqValidationError is the validation error returned
// by GetAnalyticsDownloadTokenReq.Validate if the designated constraints
// aren't met.
type GetAnalyticsDownloadTokenReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetAnalyticsDownloadTokenReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetAnalyticsDownloadTokenReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetAnalyticsDownloadTokenReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetAnalyticsDownloadTokenReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetAnalyticsDownloadTokenReqValidationError) ErrorName() string {
	return "GetAnalyticsDownloadTokenReqValidationError"
}

// Error satisfies the builtin error interface
func (e GetAnalyticsDownloadTokenReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetAnalyticsDownloadTokenReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetAnalyticsDownloadTokenReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetAnalyticsDownloadTokenReqValidationError{}

// Validate checks the field values on GetAnalyticsDownloadTokenRes with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetAnalyticsDownloadTokenRes) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetAnalyticsDownloadTokenRes with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetAnalyticsDownloadTokenResMultiError, or nil if none found.
func (m *GetAnalyticsDownloadTokenRes) ValidateAll() error {
	return m.validate(true)
}

func (m *GetAnalyticsDownloadTokenRes) validate(all bool) error {
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
		return GetAnalyticsDownloadTokenResMultiError(errors)
	}

	return nil
}

// GetAnalyticsDownloadTokenResMultiError is an error wrapping multiple
// validation errors returned by GetAnalyticsDownloadTokenRes.ValidateAll() if
// the designated constraints aren't met.
type GetAnalyticsDownloadTokenResMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetAnalyticsDownloadTokenResMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetAnalyticsDownloadTokenResMultiError) AllErrors() []error { return m }

// GetAnalyticsDownloadTokenResValidationError is the validation error returned
// by GetAnalyticsDownloadTokenRes.Validate if the designated constraints
// aren't met.
type GetAnalyticsDownloadTokenResValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetAnalyticsDownloadTokenResValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetAnalyticsDownloadTokenResValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetAnalyticsDownloadTokenResValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetAnalyticsDownloadTokenResValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetAnalyticsDownloadTokenResValidationError) ErrorName() string {
	return "GetAnalyticsDownloadTokenResValidationError"
}

// Error satisfies the builtin error interface
func (e GetAnalyticsDownloadTokenResValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetAnalyticsDownloadTokenRes.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetAnalyticsDownloadTokenResValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetAnalyticsDownloadTokenResValidationError{}