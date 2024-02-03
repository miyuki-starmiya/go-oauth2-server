package model

type AuthorizationData struct {
	ClientID          string `bson:"client_id"`
	RedirectURI       string `bson:"redirect_uri"`
	AuthorizationCode string `bson:"authorization_code"`
}