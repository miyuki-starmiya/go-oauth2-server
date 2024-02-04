package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/miyuki-starmiya/go-oauth2-server/db/store"
	"github.com/miyuki-starmiya/go-oauth2-server/resource/domain/repository"
)

func NewResourceHandler(ts *store.TokenStore) *ResourceHandler {
	return &ResourceHandler{
		TokenStore: ts,
	}
}

type ResourceHandler struct {
	TokenStore *store.TokenStore
}

func (rh *ResourceHandler) HandleResourceRequest(w http.ResponseWriter, r *http.Request) {
	if !rh.validateAccessToken(r) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	resource, _ := repository.GetResource()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resource)
	log.Println("get resource successfully")
}

func (rh *ResourceHandler) validateAccessToken(r *http.Request) bool {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		log.Println("Authorization header is not set")
		return false
	}
	_, token, err := parseBearerToken(authHeader)
	if err != nil {
		log.Printf("error: %v\n", err)
		return false
	}

	tokenData, err := rh.TokenStore.GetData(r.URL.Query().Get("client_id"), token)
	if err != nil {
		log.Printf("Client ID and access token do not match: %v\n", err)
		return false
	}
	if tokenData.IssuedAt.Add(time.Duration(tokenData.ExpiresIn) * time.Second).Before(time.Now()) {
		log.Println("Access token has expired")
		return false
	}

	return true
}

func parseBearerToken(authHeader string) (string, string, error) {
	const bearerPrefix = "Bearer "
	if len(authHeader) < len(bearerPrefix) || authHeader[:len(bearerPrefix)] != bearerPrefix {
		return "", "", fmt.Errorf("Invalid authorization header")
	}

	return "Bearer", authHeader[len(bearerPrefix):], nil
}
