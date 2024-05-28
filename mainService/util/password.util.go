package util

import (
	"crypto/sha256"
	"encoding/hex"
)

// Sha256 computes the SHA-256 digest of an input string and
// returns a hex encoded representation of the digest.
func Sha256(password string) string {
	digest := sha256.Sum256([]byte(password))
	return hex.EncodeToString(digest[:])
}
