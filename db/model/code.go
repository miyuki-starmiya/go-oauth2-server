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
