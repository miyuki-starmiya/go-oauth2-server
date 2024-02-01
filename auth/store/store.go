package store

import (
	"fmt"

	"go-oauth2-server/auth/model"
)

func NewStore() *Store {
	return &Store{}
}

type Store struct {
	Data []*model.AuthorizationData
}

func (s *Store) GetData(clientId string, authorizationCode string) (*model.AuthorizationData, error) {
	authorizationData := &model.AuthorizationData{}
	for _, data := range s.Data {
		if (data.ClientID == clientId) && (data.AuthorizationCode == authorizationCode) {
			authorizationData = data
		} else {
			return nil, fmt.Errorf("authorization code with this client is not found")
		}
	}
	return authorizationData, nil
}

func (s *Store) SetData(data *model.AuthorizationData) error {
	s.Data = append(s.Data, data)
	return nil
}
