// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: plugnmeet_nats_msg.proto

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

// Validate checks the field values on NatsMsgServerToClient with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *NatsMsgServerToClient) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on NatsMsgServerToClient with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// NatsMsgServerToClientMultiError, or nil if none found.
func (m *NatsMsgServerToClient) ValidateAll() error {
	return m.validate(true)
}

func (m *NatsMsgServerToClient) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Event

	// no validation rules for Msg

	if len(errors) > 0 {
		return NatsMsgServerToClientMultiError(errors)
	}

	return nil
}

// NatsMsgServerToClientMultiError is an error wrapping multiple validation
// errors returned by NatsMsgServerToClient.ValidateAll() if the designated
// constraints aren't met.
type NatsMsgServerToClientMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m NatsMsgServerToClientMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m NatsMsgServerToClientMultiError) AllErrors() []error { return m }

// NatsMsgServerToClientValidationError is the validation error returned by
// NatsMsgServerToClient.Validate if the designated constraints aren't met.
type NatsMsgServerToClientValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NatsMsgServerToClientValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NatsMsgServerToClientValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NatsMsgServerToClientValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NatsMsgServerToClientValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NatsMsgServerToClientValidationError) ErrorName() string {
	return "NatsMsgServerToClientValidationError"
}

// Error satisfies the builtin error interface
func (e NatsMsgServerToClientValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNatsMsgServerToClient.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NatsMsgServerToClientValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NatsMsgServerToClientValidationError{}

// Validate checks the field values on NatsMsgClientToServer with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *NatsMsgClientToServer) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on NatsMsgClientToServer with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// NatsMsgClientToServerMultiError, or nil if none found.
func (m *NatsMsgClientToServer) ValidateAll() error {
	return m.validate(true)
}

func (m *NatsMsgClientToServer) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Event

	// no validation rules for Msg

	if len(errors) > 0 {
		return NatsMsgClientToServerMultiError(errors)
	}

	return nil
}

// NatsMsgClientToServerMultiError is an error wrapping multiple validation
// errors returned by NatsMsgClientToServer.ValidateAll() if the designated
// constraints aren't met.
type NatsMsgClientToServerMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m NatsMsgClientToServerMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m NatsMsgClientToServerMultiError) AllErrors() []error { return m }

// NatsMsgClientToServerValidationError is the validation error returned by
// NatsMsgClientToServer.Validate if the designated constraints aren't met.
type NatsMsgClientToServerValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NatsMsgClientToServerValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NatsMsgClientToServerValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NatsMsgClientToServerValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NatsMsgClientToServerValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NatsMsgClientToServerValidationError) ErrorName() string {
	return "NatsMsgClientToServerValidationError"
}

// Error satisfies the builtin error interface
func (e NatsMsgClientToServerValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNatsMsgClientToServer.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NatsMsgClientToServerValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NatsMsgClientToServerValidationError{}

// Validate checks the field values on NatsKvRoomInfo with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *NatsKvRoomInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on NatsKvRoomInfo with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in NatsKvRoomInfoMultiError,
// or nil if none found.
func (m *NatsKvRoomInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *NatsKvRoomInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for RoomId

	// no validation rules for RoomSid

	// no validation rules for Metadata

	if len(errors) > 0 {
		return NatsKvRoomInfoMultiError(errors)
	}

	return nil
}

// NatsKvRoomInfoMultiError is an error wrapping multiple validation errors
// returned by NatsKvRoomInfo.ValidateAll() if the designated constraints
// aren't met.
type NatsKvRoomInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m NatsKvRoomInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m NatsKvRoomInfoMultiError) AllErrors() []error { return m }

// NatsKvRoomInfoValidationError is the validation error returned by
// NatsKvRoomInfo.Validate if the designated constraints aren't met.
type NatsKvRoomInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NatsKvRoomInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NatsKvRoomInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NatsKvRoomInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NatsKvRoomInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NatsKvRoomInfoValidationError) ErrorName() string { return "NatsKvRoomInfoValidationError" }

// Error satisfies the builtin error interface
func (e NatsKvRoomInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNatsKvRoomInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NatsKvRoomInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NatsKvRoomInfoValidationError{}

// Validate checks the field values on NatsKvUserInfo with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *NatsKvUserInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on NatsKvUserInfo with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in NatsKvUserInfoMultiError,
// or nil if none found.
func (m *NatsKvUserInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *NatsKvUserInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for UserId

	// no validation rules for UserSid

	// no validation rules for Name

	// no validation rules for RoomId

	// no validation rules for Metadata

	if len(errors) > 0 {
		return NatsKvUserInfoMultiError(errors)
	}

	return nil
}

// NatsKvUserInfoMultiError is an error wrapping multiple validation errors
// returned by NatsKvUserInfo.ValidateAll() if the designated constraints
// aren't met.
type NatsKvUserInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m NatsKvUserInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m NatsKvUserInfoMultiError) AllErrors() []error { return m }

// NatsKvUserInfoValidationError is the validation error returned by
// NatsKvUserInfo.Validate if the designated constraints aren't met.
type NatsKvUserInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NatsKvUserInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NatsKvUserInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NatsKvUserInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NatsKvUserInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NatsKvUserInfoValidationError) ErrorName() string { return "NatsKvUserInfoValidationError" }

// Error satisfies the builtin error interface
func (e NatsKvUserInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNatsKvUserInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NatsKvUserInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NatsKvUserInfoValidationError{}
