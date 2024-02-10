package model

import (
	"github.com/miyuki-starmiya/go-oauth2-server/db/constants"
)

type AuthorizationData struct {
	ClientID            string                         `bson:"client_id"`
	RedirectURI         string                         `bson:"redirect_uri"`
	AuthorizationCode   string                         `bson:"authorization_code"`
	CodeChallenge       *string                        `bson:"code_challenge"`
	CodeChallengeMethod *constants.CodeChallengeMethod `bson:"code_challenge_method"`
}

type AuthorizationDataOption func(*AuthorizationData)

func WithCodeChallenge(codeChallenge string) AuthorizationDataOption {
	return func(ad *AuthorizationData) {
		ad.CodeChallenge = &codeChallenge
	}
}

func WithCodeChallengeMethod(codeChallengeMethod constants.CodeChallengeMethod) AuthorizationDataOption {
	return func(ad *AuthorizationData) {
		ad.CodeChallengeMethod = &codeChallengeMethod
	}
}

func NewAuthorizationData(clientID, redirectURI, authorizationCode string, opts ...AuthorizationDataOption) *AuthorizationData {
	ad := &AuthorizationData{
		ClientID:          clientID,
		RedirectURI:       redirectURI,
		AuthorizationCode: authorizationCode,
	}
	for _, opt := range opts {
		opt(ad)
	}

	return ad
}
