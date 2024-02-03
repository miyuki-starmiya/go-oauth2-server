package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"go-oauth2-server/auth/generate"
	"go-oauth2-server/auth/model"
	"go-oauth2-server/auth/store"
	"go-oauth2-server/auth/util"
)

func NewTokenHandler(cs *store.CodeStore, ts *store.TokenStore) *TokenHandler {
	return &TokenHandler{
		CodeStore:  cs,
		TokenStore: ts,
	}
}

type TokenHandler struct {
	CodeStore  *store.CodeStore
	TokenStore *store.TokenStore
}

func (th *TokenHandler) HandleTokenRequest(w http.ResponseWriter, r *http.Request) {
	if !th.validateTokenRequest(r) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	clientId, _, _ := getClientData(r)
	access, refresh, _ := generate.NewAccessGenerate().Token(r.Context(), clientId, false)
	tokenType := "Bearer"
	expiresIn := 3600

	// store the token object
	tokenData := &model.TokenData{
		ClientID:     clientId,
		AccessToken:  access,
		ExpiresIn:    expiresIn,
		RefreshToken: refresh,
		IssuedAt:     time.Now(),
		TokenType:    tokenType,
	}
	if err := th.TokenStore.CreateData(tokenData); err != nil {
		log.Printf("Error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"access_token":  access,
		"token_type":    tokenType,
		"expires_in":    expiresIn,
		"refresh_token": refresh,
	})
	log.Println("access token generated successfully")
}

func (th *TokenHandler) validateTokenRequest(r *http.Request) bool {
	// Check Request Parameters
	if r.Method != "POST" {
		log.Println("request method must be POST")
		return false
	}
	if r.URL.Query().Get("grant_type") != "authorization_code" {
		log.Println("grant_type must be authorization_code")
		return false
	}
	if r.URL.Query().Get("redirect_uri") != os.Getenv("REDIRECT_URI") {
		log.Println("redirect_uri is wrong")
		return false
	}

	// Check Basic Authentication Header
	clientId, clientSecret, err := getClientData(r)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return false
	}
	if _, err := th.CodeStore.GetData(clientId, r.URL.Query().Get("code")); err != nil {
		log.Printf("Error: %v\n", err)
		return false
	}
	if clientSecret != os.Getenv("CLIENT_SECRET") {
		log.Println("client secret is wrong")
		return false
	}

	return true
}

func getClientData(r *http.Request) (string, string, error) {
	authorizationHeader, err := util.RetrieveAuthorizationHeader(r)
	if err != nil {
		return "", "", err
	}
	clientId, clientSecret, err := util.DecodeClientBase64(authorizationHeader)
	if err != nil {
		return "", "", err
	}

	return clientId, clientSecret, nil
}
