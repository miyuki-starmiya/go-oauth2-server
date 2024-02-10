package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/miyuki-starmiya/go-oauth2-server/auth/generate"
	"github.com/miyuki-starmiya/go-oauth2-server/auth/util"
	"github.com/miyuki-starmiya/go-oauth2-server/db/model"
	"github.com/miyuki-starmiya/go-oauth2-server/db/store"
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

type TokenRequest struct {
	GrantType    string  `json:"grant_type"`
	RedirectURI  string  `json:"redirect_uri"`
	Code         string  `json:"code"`
	CodeVerifier *string `json:"code_verifier"`
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
	// extract request body
	var tr TokenRequest
	if err := json.NewDecoder(r.Body).Decode(&tr); err != nil {
		log.Printf("Error: %v\n", err)
		return false
	}
	defer r.Body.Close()
	log.Println("got token request:", tr)

	if r.Method != "POST" {
		log.Println("request method must be POST")
		return false
	}
	if tr.GrantType != "authorization_code" {
		log.Println("grant_type must be authorization_code")
		return false
	}
	if tr.RedirectURI != os.Getenv("REDIRECT_URI") {
		log.Println("redirect_uri is wrong")
		return false
	}

	// Check Basic Authentication Header
	clientId, clientSecret, err := getClientData(r)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return false
	}
	if clientSecret != os.Getenv("CLIENT_SECRET") {
		log.Println("client secret is wrong")
		return false
	}
	authorizationData, err := th.CodeStore.GetData(clientId, tr.Code)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return false
	}
	if validatePKCETokenRequest(&tr, authorizationData) == false {
		return false
	}

	return true
}

func validatePKCETokenRequest(tr *TokenRequest, ad *model.AuthorizationData) bool {
	log.Println("got authorizationData:", ad)
	if ad.CodeChallenge == nil && ad.CodeChallengeMethod == nil {
		return true
	}

	if *tr.CodeVerifier == "" {
		log.Println("code_verifier is empty")
		return false
	} else if *tr.CodeVerifier != "" && util.GenerateCodeChallenge(*tr.CodeVerifier, *ad.CodeChallengeMethod) != *ad.CodeChallenge {
		log.Println("code_verifier is wrong")
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
