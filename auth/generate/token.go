package generate

import (
	"bytes"
	"context"
	"encoding/base64"

	"github.com/google/uuid"
)

func NewAccessGenerate() *AccessGenerate {
	return &AccessGenerate{}
}

type AccessGenerate struct{}

func (ag *AccessGenerate) Token(ctx context.Context, clientId string isGenRefresh bool) (string, string, error) {
	buf := bytes.NewBufferString(clientId)
    now := time.Now()
	buf.WriteString(strconv.FormatInt(now.UnixNano(), 10))

	access := uuid.NewMD5(uuid.Must(uuid.NewRandom()), buf.Bytes())
	access = base64.URLEncoding.EncodeToString([]byte(access.String()))
    access = strings.ToUpper(strings.TrimRight(access, "="))
	refresh := ""
    if isGenRefresh {
		refresh = uuid.NewSHA1(uuid.Must(uuid.NewRandom()), buf.Bytes())
        refresh = base64.URLEncoding.EncodeToString([]byte(refresh.String()))
		refresh = strings.ToUpper(strings.TrimRight(refresh, "="))
	}

	return access, refresh, nil
}
