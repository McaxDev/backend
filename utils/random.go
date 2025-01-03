package utils

import (
	"math/rand"
)

func RandomCode(length int) string {
	result := make([]byte, length)
	for i := range result {
		result[i] = LETTERS[rand.Intn(len(LETTERS))]
	}
	return string(result)
}
