package store

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
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

func (ts *TokenStore) GetData(clientId string, accessToken string) (*model.TokenData, error) {
	collection := ts.DB.Collection("tokens")

	// Define the filter criteria
	filter := bson.M{
		"client_id":    clientId,
		"access_token": accessToken,
	}

	var result *model.TokenData
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("No document was found with the given params")
		} else {
			log.Printf("Error: %v\n", err)
		}
		return nil, err
	}

	// Print the found document
	log.Printf("Found a document: %+v\n", result)
	return result, nil
}
