package model

type AuthorizationData struct {
	ClientID          string `json:"client_id"`
	RedirectURI       string `json:"redirect_uri"`
	AuthorizationCode string `json:"authorization_code"`
}
