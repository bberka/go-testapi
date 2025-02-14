package utils

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

// HashMD5 computes the MD5 hash of the given plain text and returns it as a hexadecimal string.
func HashMD5(plainText string) string {
	hasher := md5.New()
	hasher.Write([]byte(plainText))
	return hex.EncodeToString(hasher.Sum(nil))
}

// HashSHA256 computes the SHA-256 hash of the given plain text and returns it as a hexadecimal string.
func HashSHA256(plainText string) string {
	hasher := sha256.New()
	hasher.Write([]byte(plainText))
	return hex.EncodeToString(hasher.Sum(nil))
}

func HashPassword(plainText string) string {
	return HashMD5(HashSHA256(plainText))
}
