package utils

import (
	"fmt"
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateShortCode(n int) string {
	seed := time.Now().UnixNano()
	src := rand.NewSource(seed)
	random := rand.New(src)

	fmt.Println("random no:", random)
	fmt.Println("random int:", random.Intn(100))

	b := make([]byte, n)

	fmt.Println("b:", b)

	for i := range b {
		b[i] = charset[random.Intn(len(charset))]
	}

	fmt.Println("b", string(b))

	return string(b)
}
