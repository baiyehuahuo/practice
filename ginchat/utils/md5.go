package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// Md5EncodeSmall small encode bytes
func Md5EncodeSmall(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// Md5EncodeBig big encode bytes
func Md5EncodeBig(str string) string {
	return strings.ToUpper(Md5EncodeSmall(str))
}

// MakePassword encoding
func MakePassword(password, salt string) string {
	return Md5EncodeSmall(password + salt)
}

// ValidPassword decoding
func ValidPassword(input, salt, password string) bool {
	return Md5EncodeSmall(input+salt) == password
}
