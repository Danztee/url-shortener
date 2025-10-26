package utils

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateShortCode(n int) string {
	seed := time.Now().UnixNano()
	src := rand.NewSource(seed)
	random := rand.New(src)

	b := make([]byte, n)

	for i := range b {
		b[i] = charset[random.Intn(len(charset))]
	}

	return string(b)
}
