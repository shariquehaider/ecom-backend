package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var JwtSecret = []byte("my_secret_key")

func GenerateJWT(id string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"_id": id,
		"exp": time.Now().Add(time.Hour * 72).Unix(), // Token expires in 72 hours
	})

	tokenString, err := token.SignedString(JwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
