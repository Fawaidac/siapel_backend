package utils

import (
	"math/rand"
	"time"
)

func GenerateTicketString(prefix string) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := make([]byte, 6)
	for i := range result {
		result[i] = charset[seededRand.Intn(len(charset))]
	}
	return prefix + "-" + string(result)
}

