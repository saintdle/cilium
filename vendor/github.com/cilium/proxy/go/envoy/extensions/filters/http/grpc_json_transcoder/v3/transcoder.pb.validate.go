// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/filters/http/grpc_json_transcoder/v3/transcoder.proto

package grpc_json_transcoderv3

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

// Validate checks the field values on GrpcJsonTranscoder with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GrpcJsonTranscoder) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GrpcJsonTranscoder with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GrpcJsonTranscoderMultiError, or nil if none found.
func (m *GrpcJsonTranscoder) ValidateAll() error {
	return m.validate(true)
}

func (m *GrpcJsonTranscoder) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetPrintOptions()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GrpcJsonTranscoderValidationError{
					field:  "PrintOptions",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GrpcJsonTranscoderValidationError{
					field:  "PrintOptions",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPrintOptions()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GrpcJsonTranscoderValidationError{
				field:  "PrintOptions",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for MatchIncomingRequestRoute

	// no validation rules for AutoMapping

	// no validation rules for IgnoreUnknownQueryParameters

	// no validation rules for ConvertGrpcStatus

	if _, ok := GrpcJsonTranscoder_UrlUnescapeSpec_name[int32(m.GetUrlUnescapeSpec())]; !ok {
		err := GrpcJsonTranscoderValidationError{
			field:  "UrlUnescapeSpec",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for QueryParamUnescapePlus

	// no validation rules for MatchUnregisteredCustomVerb

	if all {
		switch v := interface{}(m.GetRequestValidationOptions()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GrpcJsonTranscoderValidationError{
					field:  "RequestValidationOptions",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GrpcJsonTranscoderValidationError{
					field:  "RequestValidationOptions",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRequestValidationOptions()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GrpcJsonTranscoderValidationError{
				field:  "RequestValidationOptions",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for CaseInsensitiveEnumParsing

	if wrapper := m.GetMaxRequestBodySize(); wrapper != nil {

		if wrapper.GetValue() <= 0 {
			err := GrpcJsonTranscoderValidationError{
				field:  "MaxRequestBodySize",
				reason: "value must be greater than 0",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if wrapper := m.GetMaxResponseBodySize(); wrapper != nil {

		if wrapper.GetValue() <= 0 {
			err := GrpcJsonTranscoderValidationError{
				field:  "MaxResponseBodySize",
				reason: "value must be greater than 0",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	switch m.DescriptorSet.(type) {

	case *GrpcJsonTranscoder_ProtoDescriptor:
		// no validation rules for ProtoDescriptor

	case *GrpcJsonTranscoder_ProtoDescriptorBin:
		// no validation rules for ProtoDescriptorBin

	default:
		err := GrpcJsonTranscoderValidationError{
			field:  "DescriptorSet",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)

	}

	if len(errors) > 0 {
		return GrpcJsonTranscoderMultiError(errors)
	}
	return nil
}

// GrpcJsonTranscoderMultiError is an error wrapping multiple validation errors
// returned by GrpcJsonTranscoder.ValidateAll() if the designated constraints
// aren't met.
type GrpcJsonTranscoderMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GrpcJsonTranscoderMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GrpcJsonTranscoderMultiError) AllErrors() []error { return m }

// GrpcJsonTranscoderValidationError is the validation error returned by
// GrpcJsonTranscoder.Validate if the designated constraints aren't met.
type GrpcJsonTranscoderValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GrpcJsonTranscoderValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GrpcJsonTranscoderValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GrpcJsonTranscoderValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GrpcJsonTranscoderValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GrpcJsonTranscoderValidationError) ErrorName() string {
	return "GrpcJsonTranscoderValidationError"
}

// Error satisfies the builtin error interface
func (e GrpcJsonTranscoderValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGrpcJsonTranscoder.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GrpcJsonTranscoderValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GrpcJsonTranscoderValidationError{}

// Validate checks the field values on GrpcJsonTranscoder_PrintOptions with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GrpcJsonTranscoder_PrintOptions) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GrpcJsonTranscoder_PrintOptions with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// GrpcJsonTranscoder_PrintOptionsMultiError, or nil if none found.
func (m *GrpcJsonTranscoder_PrintOptions) ValidateAll() error {
	return m.validate(true)
}

func (m *GrpcJsonTranscoder_PrintOptions) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for AddWhitespace

	// no validation rules for AlwaysPrintPrimitiveFields

	// no validation rules for AlwaysPrintEnumsAsInts

	// no validation rules for PreserveProtoFieldNames

	// no validation rules for StreamNewlineDelimited

	if len(errors) > 0 {
		return GrpcJsonTranscoder_PrintOptionsMultiError(errors)
	}
	return nil
}

// GrpcJsonTranscoder_PrintOptionsMultiError is an error wrapping multiple
// validation errors returned by GrpcJsonTranscoder_PrintOptions.ValidateAll()
// if the designated constraints aren't met.
type GrpcJsonTranscoder_PrintOptionsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GrpcJsonTranscoder_PrintOptionsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GrpcJsonTranscoder_PrintOptionsMultiError) AllErrors() []error { return m }

// GrpcJsonTranscoder_PrintOptionsValidationError is the validation error
// returned by GrpcJsonTranscoder_PrintOptions.Validate if the designated
// constraints aren't met.
type GrpcJsonTranscoder_PrintOptionsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GrpcJsonTranscoder_PrintOptionsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GrpcJsonTranscoder_PrintOptionsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GrpcJsonTranscoder_PrintOptionsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GrpcJsonTranscoder_PrintOptionsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GrpcJsonTranscoder_PrintOptionsValidationError) ErrorName() string {
	return "GrpcJsonTranscoder_PrintOptionsValidationError"
}

// Error satisfies the builtin error interface
func (e GrpcJsonTranscoder_PrintOptionsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGrpcJsonTranscoder_PrintOptions.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GrpcJsonTranscoder_PrintOptionsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GrpcJsonTranscoder_PrintOptionsValidationError{}

// Validate checks the field values on
// GrpcJsonTranscoder_RequestValidationOptions with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GrpcJsonTranscoder_RequestValidationOptions) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// GrpcJsonTranscoder_RequestValidationOptions with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in
// GrpcJsonTranscoder_RequestValidationOptionsMultiError, or nil if none found.
func (m *GrpcJsonTranscoder_RequestValidationOptions) ValidateAll() error {
	return m.validate(true)
}

func (m *GrpcJsonTranscoder_RequestValidationOptions) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for RejectUnknownMethod

	// no validation rules for RejectUnknownQueryParameters

	// no validation rules for RejectBindingBodyFieldCollisions

	if len(errors) > 0 {
		return GrpcJsonTranscoder_RequestValidationOptionsMultiError(errors)
	}
	return nil
}

// GrpcJsonTranscoder_RequestValidationOptionsMultiError is an error wrapping
// multiple validation errors returned by
// GrpcJsonTranscoder_RequestValidationOptions.ValidateAll() if the designated
// constraints aren't met.
type GrpcJsonTranscoder_RequestValidationOptionsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GrpcJsonTranscoder_RequestValidationOptionsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GrpcJsonTranscoder_RequestValidationOptionsMultiError) AllErrors() []error { return m }

// GrpcJsonTranscoder_RequestValidationOptionsValidationError is the validation
// error returned by GrpcJsonTranscoder_RequestValidationOptions.Validate if
// the designated constraints aren't met.
type GrpcJsonTranscoder_RequestValidationOptionsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GrpcJsonTranscoder_RequestValidationOptionsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GrpcJsonTranscoder_RequestValidationOptionsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GrpcJsonTranscoder_RequestValidationOptionsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GrpcJsonTranscoder_RequestValidationOptionsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GrpcJsonTranscoder_RequestValidationOptionsValidationError) ErrorName() string {
	return "GrpcJsonTranscoder_RequestValidationOptionsValidationError"
}

// Error satisfies the builtin error interface
func (e GrpcJsonTranscoder_RequestValidationOptionsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGrpcJsonTranscoder_RequestValidationOptions.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GrpcJsonTranscoder_RequestValidationOptionsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GrpcJsonTranscoder_RequestValidationOptionsValidationError{}