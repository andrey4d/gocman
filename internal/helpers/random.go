/*
 *   Copyright (c) 2023 Andrey Danilov andrey4d.dev@gmail.com
 *   All rights reserved.
 */
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
