// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/retry/previous_priorities/previous_priorities_config.proto

package previous_priorities

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

// Validate checks the field values on PreviousPrioritiesConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *PreviousPrioritiesConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PreviousPrioritiesConfig with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PreviousPrioritiesConfigMultiError, or nil if none found.
func (m *PreviousPrioritiesConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *PreviousPrioritiesConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetUpdateFrequency() <= 0 {
		err := PreviousPrioritiesConfigValidationError{
			field:  "UpdateFrequency",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return PreviousPrioritiesConfigMultiError(errors)
	}
	return nil
}

// PreviousPrioritiesConfigMultiError is an error wrapping multiple validation
// errors returned by PreviousPrioritiesConfig.ValidateAll() if the designated
// constraints aren't met.
type PreviousPrioritiesConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PreviousPrioritiesConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PreviousPrioritiesConfigMultiError) AllErrors() []error { return m }

// PreviousPrioritiesConfigValidationError is the validation error returned by
// PreviousPrioritiesConfig.Validate if the designated constraints aren't met.
type PreviousPrioritiesConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PreviousPrioritiesConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PreviousPrioritiesConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PreviousPrioritiesConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PreviousPrioritiesConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PreviousPrioritiesConfigValidationError) ErrorName() string {
	return "PreviousPrioritiesConfigValidationError"
}

// Error satisfies the builtin error interface
func (e PreviousPrioritiesConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPreviousPrioritiesConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PreviousPrioritiesConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PreviousPrioritiesConfigValidationError{}
