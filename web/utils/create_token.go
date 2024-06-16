package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

func CreateToken(username string, email string) (string, error) {

	claims := jwt.MapClaims{
		"username": username,
		"email":    email,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		return "", err
	}
	return tokenString, nil
}
