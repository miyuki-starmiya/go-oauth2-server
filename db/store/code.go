package store

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"go-oauth2-server/auth/model"
)

func NewCodeStore(db *mongo.Database) *CodeStore {
	return &CodeStore{
		DB: db,
	}
}

type CodeStore struct {
	DB *mongo.Database
}

func (cs *CodeStore) CreateData(data *model.AuthorizationData) error {
	collection := cs.DB.Collection("codes")

	// Insert a single document
	insertResult, err := collection.InsertOne(context.TODO(), data)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return err
	}
	log.Printf("Inserted a single document: %v\n", insertResult.InsertedID)
	return nil
}

func (cs *CodeStore) GetData(clientId string, authorizationCode string) (*model.AuthorizationData, error) {
	collection := cs.DB.Collection("codes")

	// Define the filter criteria
	filter := bson.M{
		"client_id":          clientId,
		"authorization_code": authorizationCode,
	}

	var result *model.AuthorizationData
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
