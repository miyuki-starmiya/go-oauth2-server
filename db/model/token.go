package model

import "time"

type TokenData struct {
	ClientID     string    `bson:"client_id"`
	AccessToken  string    `bson:"access_token"`
	IssuedAt     time.Time `bson:"issued_at"`
	ExpiresIn    int       `bson:"expires_in"`
	RefreshToken string    `bson:"refresh_token"`
	TokenType    string    `bson:"token_type"`
}
