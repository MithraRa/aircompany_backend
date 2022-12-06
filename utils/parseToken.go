package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func ParseToken(cookie *http.Cookie) (jwt.Token, jwt.MapClaims, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("can't parse the token")
	}

	secretKey := os.Getenv("secret_key")

	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(cookie.Value, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return *token, claims, err
	}

	return *token, claims, nil
}
