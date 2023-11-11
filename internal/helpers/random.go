package helpers

import (
	"crypto/rand"
	"encoding/hex"
)

// NewRandomString generates random string with given size.

func CreateContainerID(size int) string {
	randBytes := make([]byte, size)
	rand.Read(randBytes)
	return hex.EncodeToString(randBytes)
}
