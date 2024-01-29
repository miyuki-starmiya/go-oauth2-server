package util

import (
	"fmt"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func GetAccessToken(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", fmt.Errorf("Authorization header is not set")
	}

	_, token, err := parseBearerToken(authHeader)
	if err != nil {
		return "", err
	}

	if !validateAccessToken(token) {
		return "", fmt.Errorf("Invalid access token")
	}

	return token, nil
}

func parseBearerToken(authHeader string) (string, string, error) {
	const bearerPrefix = "Bearer "
	if len(authHeader) < len(bearerPrefix) || authHeader[:len(bearerPrefix)] != bearerPrefix {
		return "", "", fmt.Errorf("Invalid authorization header")
	}

	return "Bearer", authHeader[len(bearerPrefix):], nil
}

func validateAccessToken(token string) bool {
	accessToken := os.Getenv("ACCESS_TOKEN")
	if accessToken == token {
		return true
	} else {
		return false
	}
}
