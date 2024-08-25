// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: livekit_agent_dispatch.proto

package livekit

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

// Validate checks the field values on CreateAgentDispatchRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateAgentDispatchRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateAgentDispatchRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateAgentDispatchRequestMultiError, or nil if none found.
func (m *CreateAgentDispatchRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateAgentDispatchRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for AgentName

	// no validation rules for Room

	// no validation rules for Metadata

	if len(errors) > 0 {
		return CreateAgentDispatchRequestMultiError(errors)
	}

	return nil
}

// CreateAgentDispatchRequestMultiError is an error wrapping multiple
// validation errors returned by CreateAgentDispatchRequest.ValidateAll() if
// the designated constraints aren't met.
type CreateAgentDispatchRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateAgentDispatchRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateAgentDispatchRequestMultiError) AllErrors() []error { return m }

// CreateAgentDispatchRequestValidationError is the validation error returned
// by CreateAgentDispatchRequest.Validate if the designated constraints aren't met.
type CreateAgentDispatchRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateAgentDispatchRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateAgentDispatchRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateAgentDispatchRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateAgentDispatchRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateAgentDispatchRequestValidationError) ErrorName() string {
	return "CreateAgentDispatchRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateAgentDispatchRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateAgentDispatchRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateAgentDispatchRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateAgentDispatchRequestValidationError{}

// Validate checks the field values on RoomAgentDispatch with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *RoomAgentDispatch) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RoomAgentDispatch with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RoomAgentDispatchMultiError, or nil if none found.
func (m *RoomAgentDispatch) ValidateAll() error {
	return m.validate(true)
}

func (m *RoomAgentDispatch) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for AgentName

	// no validation rules for Metadata

	if len(errors) > 0 {
		return RoomAgentDispatchMultiError(errors)
	}

	return nil
}

// RoomAgentDispatchMultiError is an error wrapping multiple validation errors
// returned by RoomAgentDispatch.ValidateAll() if the designated constraints
// aren't met.
type RoomAgentDispatchMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RoomAgentDispatchMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RoomAgentDispatchMultiError) AllErrors() []error { return m }

// RoomAgentDispatchValidationError is the validation error returned by
// RoomAgentDispatch.Validate if the designated constraints aren't met.
type RoomAgentDispatchValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RoomAgentDispatchValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RoomAgentDispatchValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RoomAgentDispatchValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RoomAgentDispatchValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RoomAgentDispatchValidationError) ErrorName() string {
	return "RoomAgentDispatchValidationError"
}

// Error satisfies the builtin error interface
func (e RoomAgentDispatchValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRoomAgentDispatch.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RoomAgentDispatchValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RoomAgentDispatchValidationError{}

// Validate checks the field values on DeleteAgentDispatchRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteAgentDispatchRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteAgentDispatchRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteAgentDispatchRequestMultiError, or nil if none found.
func (m *DeleteAgentDispatchRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteAgentDispatchRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for DispatchId

	// no validation rules for Room

	if len(errors) > 0 {
		return DeleteAgentDispatchRequestMultiError(errors)
	}

	return nil
}

// DeleteAgentDispatchRequestMultiError is an error wrapping multiple
// validation errors returned by DeleteAgentDispatchRequest.ValidateAll() if
// the designated constraints aren't met.
type DeleteAgentDispatchRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteAgentDispatchRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteAgentDispatchRequestMultiError) AllErrors() []error { return m }

// DeleteAgentDispatchRequestValidationError is the validation error returned
// by DeleteAgentDispatchRequest.Validate if the designated constraints aren't met.
type DeleteAgentDispatchRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteAgentDispatchRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteAgentDispatchRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteAgentDispatchRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteAgentDispatchRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteAgentDispatchRequestValidationError) ErrorName() string {
	return "DeleteAgentDispatchRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteAgentDispatchRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteAgentDispatchRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteAgentDispatchRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteAgentDispatchRequestValidationError{}

// Validate checks the field values on ListAgentDispatchRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListAgentDispatchRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListAgentDispatchRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListAgentDispatchRequestMultiError, or nil if none found.
func (m *ListAgentDispatchRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListAgentDispatchRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for DispatchId

	// no validation rules for Room

	if len(errors) > 0 {
		return ListAgentDispatchRequestMultiError(errors)
	}

	return nil
}

// ListAgentDispatchRequestMultiError is an error wrapping multiple validation
// errors returned by ListAgentDispatchRequest.ValidateAll() if the designated
// constraints aren't met.
type ListAgentDispatchRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListAgentDispatchRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListAgentDispatchRequestMultiError) AllErrors() []error { return m }

// ListAgentDispatchRequestValidationError is the validation error returned by
// ListAgentDispatchRequest.Validate if the designated constraints aren't met.
type ListAgentDispatchRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListAgentDispatchRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListAgentDispatchRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListAgentDispatchRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListAgentDispatchRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListAgentDispatchRequestValidationError) ErrorName() string {
	return "ListAgentDispatchRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListAgentDispatchRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListAgentDispatchRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListAgentDispatchRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListAgentDispatchRequestValidationError{}

// Validate checks the field values on ListAgentDispatchResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListAgentDispatchResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListAgentDispatchResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListAgentDispatchResponseMultiError, or nil if none found.
func (m *ListAgentDispatchResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListAgentDispatchResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetAgentDispatches() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListAgentDispatchResponseValidationError{
						field:  fmt.Sprintf("AgentDispatches[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListAgentDispatchResponseValidationError{
						field:  fmt.Sprintf("AgentDispatches[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListAgentDispatchResponseValidationError{
					field:  fmt.Sprintf("AgentDispatches[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListAgentDispatchResponseMultiError(errors)
	}

	return nil
}

// ListAgentDispatchResponseMultiError is an error wrapping multiple validation
// errors returned by ListAgentDispatchResponse.ValidateAll() if the
// designated constraints aren't met.
type ListAgentDispatchResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListAgentDispatchResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListAgentDispatchResponseMultiError) AllErrors() []error { return m }

// ListAgentDispatchResponseValidationError is the validation error returned by
// ListAgentDispatchResponse.Validate if the designated constraints aren't met.
type ListAgentDispatchResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListAgentDispatchResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListAgentDispatchResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListAgentDispatchResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListAgentDispatchResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListAgentDispatchResponseValidationError) ErrorName() string {
	return "ListAgentDispatchResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListAgentDispatchResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListAgentDispatchResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListAgentDispatchResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListAgentDispatchResponseValidationError{}

// Validate checks the field values on AgentDispatch with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AgentDispatch) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AgentDispatch with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AgentDispatchMultiError, or
// nil if none found.
func (m *AgentDispatch) ValidateAll() error {
	return m.validate(true)
}

func (m *AgentDispatch) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for AgentName

	// no validation rules for Room

	// no validation rules for Metadata

	if all {
		switch v := interface{}(m.GetState()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AgentDispatchValidationError{
					field:  "State",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AgentDispatchValidationError{
					field:  "State",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetState()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AgentDispatchValidationError{
				field:  "State",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return AgentDispatchMultiError(errors)
	}

	return nil
}

// AgentDispatchMultiError is an error wrapping multiple validation errors
// returned by AgentDispatch.ValidateAll() if the designated constraints
// aren't met.
type AgentDispatchMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AgentDispatchMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AgentDispatchMultiError) AllErrors() []error { return m }

// AgentDispatchValidationError is the validation error returned by
// AgentDispatch.Validate if the designated constraints aren't met.
type AgentDispatchValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AgentDispatchValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AgentDispatchValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AgentDispatchValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AgentDispatchValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AgentDispatchValidationError) ErrorName() string { return "AgentDispatchValidationError" }

// Error satisfies the builtin error interface
func (e AgentDispatchValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAgentDispatch.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AgentDispatchValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AgentDispatchValidationError{}

// Validate checks the field values on AgentDispatchState with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AgentDispatchState) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AgentDispatchState with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AgentDispatchStateMultiError, or nil if none found.
func (m *AgentDispatchState) ValidateAll() error {
	return m.validate(true)
}

func (m *AgentDispatchState) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetJobs() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, AgentDispatchStateValidationError{
						field:  fmt.Sprintf("Jobs[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, AgentDispatchStateValidationError{
						field:  fmt.Sprintf("Jobs[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AgentDispatchStateValidationError{
					field:  fmt.Sprintf("Jobs[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for CreatedAt

	// no validation rules for DeletedAt

	if len(errors) > 0 {
		return AgentDispatchStateMultiError(errors)
	}

	return nil
}

// AgentDispatchStateMultiError is an error wrapping multiple validation errors
// returned by AgentDispatchState.ValidateAll() if the designated constraints
// aren't met.
type AgentDispatchStateMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AgentDispatchStateMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AgentDispatchStateMultiError) AllErrors() []error { return m }

// AgentDispatchStateValidationError is the validation error returned by
// AgentDispatchState.Validate if the designated constraints aren't met.
type AgentDispatchStateValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AgentDispatchStateValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AgentDispatchStateValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AgentDispatchStateValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AgentDispatchStateValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AgentDispatchStateValidationError) ErrorName() string {
	return "AgentDispatchStateValidationError"
}

// Error satisfies the builtin error interface
func (e AgentDispatchStateValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAgentDispatchState.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AgentDispatchStateValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AgentDispatchStateValidationError{}
