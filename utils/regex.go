package utils

import (
	"regexp"
)

// Base64Regex is the global compiled regex for validating base64s
var Base64Regex *regexp.Regexp

// IsBase64Valid tests if the string provided is a valid base64 value or not
func IsBase64Valid(val string) bool {
	return Base64Regex.Match([]byte(val))
}
