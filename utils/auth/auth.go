package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// TODO!!! Перенести этот ключ в переменные окружения
var JwtSecret = []byte("bcnd3mfu4roej7snbc3hry2tpyi1djet")

func GenerateJWT(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString(JwtSecret)
	return t, err
}
