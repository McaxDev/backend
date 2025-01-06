package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTClaims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

func GetJwt(userID uint) (string, error) {
	return jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		JWTClaims{
			UserID: userID,
			RegisteredClaims: jwt.RegisteredClaims{
				Issuer:    "Axolotland",
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			},
		},
	).SignedString([]byte(JWTKey))
}
