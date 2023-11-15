/*
 *   Copyright (c) 2023 Andrey Danilov andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package helpers

import (
	"crypto/rand"
	"encoding/base32"
	"encoding/hex"
)

// NewRandomString generates random string with given size.

func CreateContainerID(size int) string {
	randBytes := make([]byte, size)
	rand.Read(randBytes)

	return hex.EncodeToString(randBytes)
}

func GenerateID(size int) string {
	randBytes := make([]byte, size)
	rand.Read(randBytes)
	return base32.StdEncoding.EncodeToString(randBytes)[:size]
}
