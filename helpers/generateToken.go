package helpers

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userId int) (token string, err error) {
	key := []byte(os.Getenv("JWT_TOKEN"))
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"exp": time.Now().Add(time.Minute * 1).Unix(),
	})

	token, err = t.SignedString(key)
	if err != nil {
		return "", err
	}

	return token, nil
}