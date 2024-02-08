package generate

import (
	"bytes"
	"context"
	"encoding/base64"

	"github.com/google/uuid"
)

func NewAuthorizeGenerate() *AuthorizeGenerate {
	return &AuthorizeGenerate{}
}

type AuthorizeGenerate struct{}

func (ag *AuthorizeGenerate) Token(ctx context.Context, clientId string) (string, error) {
	buf := bytes.NewBufferString(clientId)
	token := uuid.NewMD5(uuid.Must(uuid.NewRandom()), buf.Bytes())
	code := base64.URLEncoding.EncodeToString([]byte(token.String()))

	return code, nil
}
