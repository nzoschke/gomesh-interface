// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: users/v2/users.proto

package v2pb

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// Validate checks the field values on User with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *User) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Parent

	// no validation rules for Name

	// no validation rules for DisplayName

	if v, ok := interface{}(m.GetCreateTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserValidationError{
				field:  "CreateTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetWidgets() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return UserValidationError{
					field:  fmt.Sprintf("Widgets[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// UserValidationError is the validation error returned by User.Validate if the
// designated constraints aren't met.
type UserValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserValidationError) ErrorName() string { return "UserValidationError" }

// Error satisfies the builtin error interface
func (e UserValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUser.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserValidationError{}

// Validate checks the field values on GetRequest with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *GetRequest) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetName()) > 512 {
		return GetRequestValidationError{
			field:  "Name",
			reason: "value length must be at most 512 bytes",
		}
	}

	if !_GetRequest_Name_Pattern.MatchString(m.GetName()) {
		return GetRequestValidationError{
			field:  "Name",
			reason: "value does not match regex pattern \"^orgs/[a-z0-9._-]+/users/[a-z0-9._-]+$\"",
		}
	}

	return nil
}

// GetRequestValidationError is the validation error returned by
// GetRequest.Validate if the designated constraints aren't met.
type GetRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetRequestValidationError) ErrorName() string { return "GetRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetRequestValidationError{}

var _GetRequest_Name_Pattern = regexp.MustCompile("^orgs/[a-z0-9._-]+/users/[a-z0-9._-]+$")

// Validate checks the field values on CreateRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *CreateRequest) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetParent()) > 512 {
		return CreateRequestValidationError{
			field:  "Parent",
			reason: "value length must be at most 512 bytes",
		}
	}

	if !_CreateRequest_Parent_Pattern.MatchString(m.GetParent()) {
		return CreateRequestValidationError{
			field:  "Parent",
			reason: "value does not match regex pattern \"^orgs/[a-z0-9._-]+$\"",
		}
	}

	if len(m.GetUserId()) > 512 {
		return CreateRequestValidationError{
			field:  "UserId",
			reason: "value length must be at most 512 bytes",
		}
	}

	if !_CreateRequest_UserId_Pattern.MatchString(m.GetUserId()) {
		return CreateRequestValidationError{
			field:  "UserId",
			reason: "value does not match regex pattern \"^[a-z0-9._-]+$\"",
		}
	}

	if m.GetUser() == nil {
		return CreateRequestValidationError{
			field:  "User",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetUser()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateRequestValidationError{
				field:  "User",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// CreateRequestValidationError is the validation error returned by
// CreateRequest.Validate if the designated constraints aren't met.
type CreateRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateRequestValidationError) ErrorName() string { return "CreateRequestValidationError" }

// Error satisfies the builtin error interface
func (e CreateRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateRequestValidationError{}

var _CreateRequest_Parent_Pattern = regexp.MustCompile("^orgs/[a-z0-9._-]+$")

var _CreateRequest_UserId_Pattern = regexp.MustCompile("^[a-z0-9._-]+$")

// Validate checks the field values on UpdateRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *UpdateRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetUser() == nil {
		return UpdateRequestValidationError{
			field:  "User",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetUser()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateRequestValidationError{
				field:  "User",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetUpdateMask() == nil {
		return UpdateRequestValidationError{
			field:  "UpdateMask",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetUpdateMask()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateRequestValidationError{
				field:  "UpdateMask",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// UpdateRequestValidationError is the validation error returned by
// UpdateRequest.Validate if the designated constraints aren't met.
type UpdateRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateRequestValidationError) ErrorName() string { return "UpdateRequestValidationError" }

// Error satisfies the builtin error interface
func (e UpdateRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateRequestValidationError{}

// Validate checks the field values on DeleteRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *DeleteRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	return nil
}

// DeleteRequestValidationError is the validation error returned by
// DeleteRequest.Validate if the designated constraints aren't met.
type DeleteRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteRequestValidationError) ErrorName() string { return "DeleteRequestValidationError" }

// Error satisfies the builtin error interface
func (e DeleteRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteRequestValidationError{}

// Validate checks the field values on ListRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ListRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Parent

	// no validation rules for PageSize

	// no validation rules for PageToken

	return nil
}

// ListRequestValidationError is the validation error returned by
// ListRequest.Validate if the designated constraints aren't met.
type ListRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListRequestValidationError) ErrorName() string { return "ListRequestValidationError" }

// Error satisfies the builtin error interface
func (e ListRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListRequestValidationError{}

// Validate checks the field values on ListResponse with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ListResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetUsers() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListResponseValidationError{
					field:  fmt.Sprintf("Users[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for NextPageToken

	return nil
}

// ListResponseValidationError is the validation error returned by
// ListResponse.Validate if the designated constraints aren't met.
type ListResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListResponseValidationError) ErrorName() string { return "ListResponseValidationError" }

// Error satisfies the builtin error interface
func (e ListResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListResponseValidationError{}

// Validate checks the field values on BatchGetRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *BatchGetRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Parent

	return nil
}

// BatchGetRequestValidationError is the validation error returned by
// BatchGetRequest.Validate if the designated constraints aren't met.
type BatchGetRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BatchGetRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BatchGetRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BatchGetRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BatchGetRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BatchGetRequestValidationError) ErrorName() string { return "BatchGetRequestValidationError" }

// Error satisfies the builtin error interface
func (e BatchGetRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBatchGetRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BatchGetRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BatchGetRequestValidationError{}

// Validate checks the field values on BatchGetResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *BatchGetResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetUsers() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return BatchGetResponseValidationError{
					field:  fmt.Sprintf("Users[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// BatchGetResponseValidationError is the validation error returned by
// BatchGetResponse.Validate if the designated constraints aren't met.
type BatchGetResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BatchGetResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BatchGetResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BatchGetResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BatchGetResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BatchGetResponseValidationError) ErrorName() string { return "BatchGetResponseValidationError" }

// Error satisfies the builtin error interface
func (e BatchGetResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBatchGetResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BatchGetResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BatchGetResponseValidationError{}
