package utils

import (
	"crypto/rand"
	"encoding/hex"
)

func ShortID() string {
	b := make([]byte, 6)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}
