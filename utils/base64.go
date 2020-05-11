package utils

import (
	"encoding/base64"
)

// ToBase64 returns a base64 string for the byte slice payload provided
func ToBase64(payload []byte) string {
	return base64.StdEncoding.EncodeToString(payload)
}

// FromBase64 decodes a base64 string into a byte slice
func FromBase64(payload string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(payload)
}
