// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: account/v3/account_service.proto

package account

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

// define the regex for a UUID once up-front
var _account_service_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on AccountInfo with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *AccountInfo) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetAccountId() == nil {
		return AccountInfoValidationError{
			field:  "AccountId",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetAccountId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AccountInfoValidationError{
				field:  "AccountId",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for SequenceNumber

	// no validation rules for Balance

	return nil
}

// AccountInfoValidationError is the validation error returned by
// AccountInfo.Validate if the designated constraints aren't met.
type AccountInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AccountInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AccountInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AccountInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AccountInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AccountInfoValidationError) ErrorName() string { return "AccountInfoValidationError" }

// Error satisfies the builtin error interface
func (e AccountInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAccountInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AccountInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AccountInfoValidationError{}

// Validate checks the field values on CreateAccountRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateAccountRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetAccountId() == nil {
		return CreateAccountRequestValidationError{
			field:  "AccountId",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetAccountId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateAccountRequestValidationError{
				field:  "AccountId",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// CreateAccountRequestValidationError is the validation error returned by
// CreateAccountRequest.Validate if the designated constraints aren't met.
type CreateAccountRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateAccountRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateAccountRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateAccountRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateAccountRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateAccountRequestValidationError) ErrorName() string {
	return "CreateAccountRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateAccountRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateAccountRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateAccountRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateAccountRequestValidationError{}

// Validate checks the field values on CreateAccountResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateAccountResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Result

	if v, ok := interface{}(m.GetAccountInfo()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateAccountResponseValidationError{
				field:  "AccountInfo",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// CreateAccountResponseValidationError is the validation error returned by
// CreateAccountResponse.Validate if the designated constraints aren't met.
type CreateAccountResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateAccountResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateAccountResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateAccountResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateAccountResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateAccountResponseValidationError) ErrorName() string {
	return "CreateAccountResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateAccountResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateAccountResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateAccountResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateAccountResponseValidationError{}

// Validate checks the field values on GetAccountInfoRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetAccountInfoRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetAccountId() == nil {
		return GetAccountInfoRequestValidationError{
			field:  "AccountId",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetAccountId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetAccountInfoRequestValidationError{
				field:  "AccountId",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// GetAccountInfoRequestValidationError is the validation error returned by
// GetAccountInfoRequest.Validate if the designated constraints aren't met.
type GetAccountInfoRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetAccountInfoRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetAccountInfoRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetAccountInfoRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetAccountInfoRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetAccountInfoRequestValidationError) ErrorName() string {
	return "GetAccountInfoRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetAccountInfoRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetAccountInfoRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetAccountInfoRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetAccountInfoRequestValidationError{}

// Validate checks the field values on GetAccountInfoResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetAccountInfoResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Result

	if v, ok := interface{}(m.GetAccountInfo()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetAccountInfoResponseValidationError{
				field:  "AccountInfo",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// GetAccountInfoResponseValidationError is the validation error returned by
// GetAccountInfoResponse.Validate if the designated constraints aren't met.
type GetAccountInfoResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetAccountInfoResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetAccountInfoResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetAccountInfoResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetAccountInfoResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetAccountInfoResponseValidationError) ErrorName() string {
	return "GetAccountInfoResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetAccountInfoResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetAccountInfoResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetAccountInfoResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetAccountInfoResponseValidationError{}
