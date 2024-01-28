package util

import "net/http"

func GetAccessToken(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return ""
	}

	if _, token, err := parseBearerToken(authHeader); err != nil {
		return ""
	} else {
		return token
	}
}

func parseBearerToken(authHeader string) (string, string, error) {
	const bearerPrefix = "Bearer "
	if len(authHeader) < len(bearerPrefix) || authHeader[:len(bearerPrefix)] != bearerPrefix {
		return "", "", nil
	}

	return "Bearer", authHeader[len(bearerPrefix):], nil
}
