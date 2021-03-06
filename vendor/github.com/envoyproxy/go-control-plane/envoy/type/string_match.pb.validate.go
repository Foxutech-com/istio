// Code generated by protoc-gen-validate
// source: envoy/type/string_match.proto
// DO NOT EDIT!!!

package envoy_type

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

	"github.com/gogo/protobuf/types"
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
	_ = types.DynamicAny{}
)

// Validate checks the field values on StringMatch with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *StringMatch) Validate() error {
	if m == nil {
		return nil
	}

	switch m.MatchPattern.(type) {

	case *StringMatch_Simple:
		// no validation rules for Simple

	case *StringMatch_Prefix:
		// no validation rules for Prefix

	case *StringMatch_Suffix:
		// no validation rules for Suffix

	case *StringMatch_Regex:
		// no validation rules for Regex

	}

	return nil
}

// StringMatchValidationError is the validation error returned by
// StringMatch.Validate if the designated constraints aren't met.
type StringMatchValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e StringMatchValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStringMatch.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = StringMatchValidationError{}
