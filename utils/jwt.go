package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GetJwt(userID uint, jwtKey string) (string, error) {
	return jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userId": userID,
			"exp":    time.Now().Add(time.Hour * 24).Unix(),
		},
	).SignedString([]byte(jwtKey))
}
