package utils

import (
	"crypto/sha1"
	"encoding/base64"
)

// CreateSha1 return encrypted string
func CreateSha1(data []byte) string {
	_sha1 := sha1.New()
	return base64.StdEncoding.EncodeToString(_sha1.Sum(data))
}
