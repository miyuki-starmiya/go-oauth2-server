package store

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"

	"go-oauth2-server/auth/model"
)

func NewTokenStore(db *mongo.Database) *TokenStore {
	return &TokenStore{
		DB: db,
	}
}

type TokenStore struct {
	DB *mongo.Database
}

func (ts *TokenStore) CreateData(data *model.TokenData) error {
	collection := ts.DB.Collection("tokens")

	// Insert a single document
	insertResult, err := collection.InsertOne(context.TODO(), data)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return err
	}
	log.Printf("Inserted a single document: %v\n", insertResult.InsertedID)
	return nil
}
