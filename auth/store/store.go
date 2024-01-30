package store

import (
	"go-oauth2-server/auth/model"
)

func NewStore() *Store {
	return &Store{}
}

type Store struct {
	Data []*model.AuthorizeData
}

func (s *Store) GetData(authorizationCode string) (*model.AuthorizeData, error) {
	// TODO: find and error handling
	return &model.AuthorizeData{
		ClientID:          "1234",
		RedirectURI:       "http://localhost:9000/callback",
		AuthorizationCode: authorizationCode,
	}, nil
}

func (s *Store) SetData(data *model.AuthorizeData) error {
	s.Data = append(s.Data, data)
	return nil
}
