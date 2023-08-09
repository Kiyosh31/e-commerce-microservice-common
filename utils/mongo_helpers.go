package utils

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetMongoId(id string) (primitive.ObjectID, error) {
	mongoId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return primitive.NewObjectID(), fmt.Errorf("Error while parsing string to mongoId: %v", err)
	}

	return mongoId, nil
}
