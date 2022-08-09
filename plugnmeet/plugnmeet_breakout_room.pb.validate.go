// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: plugnmeet_breakout_room.proto

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

// Validate checks the field values on CreateBreakoutRoomsReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateBreakoutRoomsReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateBreakoutRoomsReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateBreakoutRoomsReqMultiError, or nil if none found.
func (m *CreateBreakoutRoomsReq) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateBreakoutRoomsReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for RoomId

	// no validation rules for RequestedUserId

	// no validation rules for Duration

	for idx, item := range m.GetRooms() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CreateBreakoutRoomsReqValidationError{
						field:  fmt.Sprintf("Rooms[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CreateBreakoutRoomsReqValidationError{
						field:  fmt.Sprintf("Rooms[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CreateBreakoutRoomsReqValidationError{
					field:  fmt.Sprintf("Rooms[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.WelcomeMsg != nil {
		// no validation rules for WelcomeMsg
	}

	if len(errors) > 0 {
		return CreateBreakoutRoomsReqMultiError(errors)
	}

	return nil
}

// CreateBreakoutRoomsReqMultiError is an error wrapping multiple validation
// errors returned by CreateBreakoutRoomsReq.ValidateAll() if the designated
// constraints aren't met.
type CreateBreakoutRoomsReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateBreakoutRoomsReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateBreakoutRoomsReqMultiError) AllErrors() []error { return m }

// CreateBreakoutRoomsReqValidationError is the validation error returned by
// CreateBreakoutRoomsReq.Validate if the designated constraints aren't met.
type CreateBreakoutRoomsReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateBreakoutRoomsReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateBreakoutRoomsReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateBreakoutRoomsReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateBreakoutRoomsReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateBreakoutRoomsReqValidationError) ErrorName() string {
	return "CreateBreakoutRoomsReqValidationError"
}

// Error satisfies the builtin error interface
func (e CreateBreakoutRoomsReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateBreakoutRoomsReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateBreakoutRoomsReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateBreakoutRoomsReqValidationError{}

// Validate checks the field values on BreakoutRoom with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *BreakoutRoom) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on BreakoutRoom with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in BreakoutRoomMultiError, or
// nil if none found.
func (m *BreakoutRoom) ValidateAll() error {
	return m.validate(true)
}

func (m *BreakoutRoom) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Title

	// no validation rules for Duration

	// no validation rules for Started

	// no validation rules for Created

	for idx, item := range m.GetUsers() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, BreakoutRoomValidationError{
						field:  fmt.Sprintf("Users[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, BreakoutRoomValidationError{
						field:  fmt.Sprintf("Users[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return BreakoutRoomValidationError{
					field:  fmt.Sprintf("Users[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return BreakoutRoomMultiError(errors)
	}

	return nil
}

// BreakoutRoomMultiError is an error wrapping multiple validation errors
// returned by BreakoutRoom.ValidateAll() if the designated constraints aren't met.
type BreakoutRoomMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BreakoutRoomMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BreakoutRoomMultiError) AllErrors() []error { return m }

// BreakoutRoomValidationError is the validation error returned by
// BreakoutRoom.Validate if the designated constraints aren't met.
type BreakoutRoomValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BreakoutRoomValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BreakoutRoomValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BreakoutRoomValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BreakoutRoomValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BreakoutRoomValidationError) ErrorName() string { return "BreakoutRoomValidationError" }

// Error satisfies the builtin error interface
func (e BreakoutRoomValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBreakoutRoom.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BreakoutRoomValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BreakoutRoomValidationError{}

// Validate checks the field values on BreakoutRoomUser with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *BreakoutRoomUser) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on BreakoutRoomUser with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// BreakoutRoomUserMultiError, or nil if none found.
func (m *BreakoutRoomUser) ValidateAll() error {
	return m.validate(true)
}

func (m *BreakoutRoomUser) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Joined

	if len(errors) > 0 {
		return BreakoutRoomUserMultiError(errors)
	}

	return nil
}

// BreakoutRoomUserMultiError is an error wrapping multiple validation errors
// returned by BreakoutRoomUser.ValidateAll() if the designated constraints
// aren't met.
type BreakoutRoomUserMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BreakoutRoomUserMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BreakoutRoomUserMultiError) AllErrors() []error { return m }

// BreakoutRoomUserValidationError is the validation error returned by
// BreakoutRoomUser.Validate if the designated constraints aren't met.
type BreakoutRoomUserValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BreakoutRoomUserValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BreakoutRoomUserValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BreakoutRoomUserValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BreakoutRoomUserValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BreakoutRoomUserValidationError) ErrorName() string { return "BreakoutRoomUserValidationError" }

// Error satisfies the builtin error interface
func (e BreakoutRoomUserValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBreakoutRoomUser.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BreakoutRoomUserValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BreakoutRoomUserValidationError{}
