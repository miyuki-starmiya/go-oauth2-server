package handler

import (
	"log"
	"net/http"
	"os"

	"github.com/miyuki-starmiya/go-oauth2-server/auth/generate"
	"github.com/miyuki-starmiya/go-oauth2-server/db/constants"
	"github.com/miyuki-starmiya/go-oauth2-server/db/model"
	"github.com/miyuki-starmiya/go-oauth2-server/db/store"
)

func NewAuthorizeHandler(cs *store.CodeStore) *AuthorizeHandler {
	return &AuthorizeHandler{
		CodeStore: cs,
	}
}

type AuthorizeHandler struct {
	CodeStore *store.CodeStore
}

func (ah *AuthorizeHandler) HandleAuthorizeRequest(w http.ResponseWriter, r *http.Request) {
	if !validateAuthorizeRequest(r) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	clientId := r.URL.Query().Get("client_id")
	redirect_uri := r.URL.Query().Get("redirect_uri")
	state := r.URL.Query().Get("state")
	code, _ := generate.NewAuthorizeGenerate().Token(r.Context(), os.Getenv("CLIENT_ID"))

	// store the code object
	authorizationData := &model.AuthorizationData{
		ClientID:          clientId,
		RedirectURI:       redirect_uri,
		AuthorizationCode: code,
	}
	if err := ah.CodeStore.CreateData(authorizationData); err != nil {
		log.Printf("Error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// redirect
	redirectURL := os.Getenv("REDIRECT_URI") + "?code=" + code + "&state=" + state
	http.Redirect(w, r, redirectURL, http.StatusFound)
}

func validateAuthorizeRequest(r *http.Request) bool {
	if r.Method != "GET" {
		log.Println("request method must be GET")
		return false
	}
	if r.URL.Query().Get("response_type") != "code" {
		log.Println("response_type must be code")
		return false
	}
	if r.URL.Query().Get("client_id") != os.Getenv("CLIENT_ID") {
		log.Println("client_id is wrong")
		return false
	}
	if r.URL.Query().Get("redirect_uri") != os.Getenv("REDIRECT_URI") {
		log.Println("redirect_uri is wrong")
		return false
	}
	if r.URL.Query().Get("state") == "" {
		log.Println("state is empty")
		return false
	}

	// PKCE
	if validatePKCEAuthorizeRequest(r) == false {
		return false
	}

	return true
}

func validatePKCEAuthorizeRequest(r *http.Request) bool {
	CodeChallenge := r.URL.Query().Get("code_challenge")
	if CodeChallenge == "" {
		log.Println("code_challenge is empty")
		return false
	} else if CodeChallenge != "" && (len(CodeChallenge) < 43 || len(CodeChallenge) > 128) {
		log.Println("code_challenge is wrong")
		return false
	}

	codeChallengeMethod := r.URL.Query().Get("code_challenge_method")
	if codeChallengeMethod == "" {
		log.Println("code_challenge_method is empty")
		return false
	} else if constants.CodeChallengeMethod(codeChallengeMethod) != constants.CodeChallengePlain && constants.CodeChallengeMethod(codeChallengeMethod) != constants.CodeChallengeS256 {
		log.Println("code_challenge_method is wrong")
		return false
	}

	return true
}
