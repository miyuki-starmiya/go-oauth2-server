package util

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
)

func RetrieveAuthorizationHeader(r *http.Request) (string, error) {
	authorizationHeader := r.Header.Get("Authorization")
	if !strings.HasPrefix(authorizationHeader, "Basic ") {
		return "", fmt.Errorf("authorization header is not basic")
	}

	return strings.Split(authorizationHeader, "Basic ")[1], nil
}

func EncodeClientBase64(clientId, clientSecret string) string {
	return base64.StdEncoding.EncodeToString([]byte(clientId + ":" + clientSecret))
}

func DecodeClientBase64(s string) (string, string, error) {
	decoded, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", "", err
	}

	clientId, clientSecret := strings.Split(string(decoded), ":")[0], strings.Split(string(decoded), ":")[1]
	return clientId, clientSecret, nil
}
