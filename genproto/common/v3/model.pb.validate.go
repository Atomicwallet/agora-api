// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: common/v3/model.proto

package common

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
var _model_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on StellarAccountId with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *StellarAccountId) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetValue()) != 56 {
		return StellarAccountIdValidationError{
			field:  "Value",
			reason: "value length must be 56 bytes",
		}
	}

	return nil
}

// StellarAccountIdValidationError is the validation error returned by
// StellarAccountId.Validate if the designated constraints aren't met.
type StellarAccountIdValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StellarAccountIdValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StellarAccountIdValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StellarAccountIdValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StellarAccountIdValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StellarAccountIdValidationError) ErrorName() string { return "StellarAccountIdValidationError" }

// Error satisfies the builtin error interface
func (e StellarAccountIdValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStellarAccountId.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StellarAccountIdValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StellarAccountIdValidationError{}

// Validate checks the field values on TransactionId with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *TransactionId) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetHash()) != 64 {
		return TransactionIdValidationError{
			field:  "Hash",
			reason: "value length must be 64 bytes",
		}
	}

	return nil
}

// TransactionIdValidationError is the validation error returned by
// TransactionId.Validate if the designated constraints aren't met.
type TransactionIdValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TransactionIdValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TransactionIdValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TransactionIdValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TransactionIdValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TransactionIdValidationError) ErrorName() string { return "TransactionIdValidationError" }

// Error satisfies the builtin error interface
func (e TransactionIdValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTransactionId.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TransactionIdValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TransactionIdValidationError{}

// Validate checks the field values on OperationId with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *OperationId) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetTransactionId() == nil {
		return OperationIdValidationError{
			field:  "TransactionId",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetTransactionId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OperationIdValidationError{
				field:  "TransactionId",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Offset

	return nil
}

// OperationIdValidationError is the validation error returned by
// OperationId.Validate if the designated constraints aren't met.
type OperationIdValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OperationIdValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OperationIdValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OperationIdValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OperationIdValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OperationIdValidationError) ErrorName() string { return "OperationIdValidationError" }

// Error satisfies the builtin error interface
func (e OperationIdValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOperationId.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OperationIdValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OperationIdValidationError{}

// Validate checks the field values on BigDecimal with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *BigDecimal) Validate() error {
	if m == nil {
		return nil
	}

	if !_BigDecimal_Value_Pattern.MatchString(m.GetValue()) {
		return BigDecimalValidationError{
			field:  "Value",
			reason: "value does not match regex pattern \"^(?:[1-9][0-9]{0,14}|0)(?:\\\\.[0-9]{0,11}[1-9])?$\"",
		}
	}

	return nil
}

// BigDecimalValidationError is the validation error returned by
// BigDecimal.Validate if the designated constraints aren't met.
type BigDecimalValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BigDecimalValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BigDecimalValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BigDecimalValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BigDecimalValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BigDecimalValidationError) ErrorName() string { return "BigDecimalValidationError" }

// Error satisfies the builtin error interface
func (e BigDecimalValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBigDecimal.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BigDecimalValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BigDecimalValidationError{}

var _BigDecimal_Value_Pattern = regexp.MustCompile("^(?:[1-9][0-9]{0,14}|0)(?:\\.[0-9]{0,11}[1-9])?$")
