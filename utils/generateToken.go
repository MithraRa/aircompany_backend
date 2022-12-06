package utils

import (
	"airline/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"time"
)

func GenereteToken(user *models.User) (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return "", err
	}

	secretKey := os.Getenv("secret_key")

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour).Unix(),
	})

	token, err := claims.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return token, err
}
