package handler

import (
	"log"
	"net/http"
	"os"

	"go-oauth2-server/auth/generate"
)

func AuthorizeHandler(w http.ResponseWriter, r *http.Request) {
	if !validateAuthorizeRequest(r) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	state := r.URL.Query().Get("state")

	// redirect
	code, _ := generate.NewAuthorizeGenerate().Token(r.Context(), os.Getenv("CLIENT_ID"))
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

	return true
}
