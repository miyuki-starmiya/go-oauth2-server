package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"go-oauth2-server/auth/store"
	"go-oauth2-server/auth/util"
)

func NewTokenHandler(store *store.Store) *TokenHandler {
	return &TokenHandler{
		Store: store,
	}
}

type TokenHandler struct {
	Store *store.Store
}

func (th *TokenHandler) HandleTokenRequest(w http.ResponseWriter, r *http.Request) {
	if !th.validateTokenRequest(r) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "success",
	})
	log.Println("success")
}

func (th *TokenHandler) validateTokenRequest(r *http.Request) bool {
	// Check Request Parameters
	if r.Method != "POST" {
		log.Println("request method must be POST")
		return false
	}
	if r.URL.Query().Get("grant_type") != "authorization_code" {
		log.Println("response_type must be authorization_code")
		return false
	}
	if r.URL.Query().Get("redirect_uri") != os.Getenv("REDIRECT_URI") {
		log.Println("redirect_uri is wrong")
		return false
	}

	// Check Basic Authentication Header
	authorizationHeader, err := util.RetrieveAuthorizationHeader(r)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return false
	}
	clientId, clientSecret, err := util.DecodeClientBase64(authorizationHeader)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return false
	}
	if _, err := th.Store.GetData(clientId, r.URL.Query().Get("code")); err != nil {
		log.Printf("Error: %v\n", err)
		return false
	}
	if clientSecret != os.Getenv("CLIENT_SECRET") {
		log.Println("client secret is wrong")
		return false
	}

	return true
}
