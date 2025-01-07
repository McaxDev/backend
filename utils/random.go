package utils

import (
	"math/rand"
	"strings"
)

func RandomCode(length uint, withLetters bool) string {
	var charset string
	if withLetters {
		charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	} else {
		charset = "0123456789"
	}

	var sb strings.Builder
	for i := uint(0); i < length; i++ {
		sb.WriteByte(charset[rand.Intn(len(charset))])
	}
	return sb.String()
}
