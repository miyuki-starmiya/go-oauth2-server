package util

import (
	"crypto/sha256"
	"encoding/base64"
	"math/rand"
	"time"

	"github.com/miyuki-starmiya/go-oauth2-server/db/constants"
)

func GenerateCodeVerifier() string {
	const length = 43
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-._~"
	// init random seed
	rand.Seed(time.Now().UnixNano())
	// create a random string
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}

	return string(b)
}

func GenerateCodeChallenge(codeVerifier string, codeChallengeMethod constants.CodeChallengeMethod) string {
	if codeChallengeMethod == constants.CodeChallengePlain {
		return codeVerifier
	}

	// Hash the code verifier using SHA-256
	h := sha256.New()
	h.Write([]byte(codeVerifier))
	hashed := h.Sum(nil)

	// Base64-url-encode the hash and remove any padding
	codeChallenge := base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(hashed)

	return codeChallenge
}
