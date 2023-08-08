package main

import (
	"context"
	"errors"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

func ConnectToDB(uri string) (*mongo.Client, error) {
	MongoClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	return MongoClient, nil
}

func GetEnvVar(envVar string) (string, error) {
	value := os.Getenv(envVar)
	if value == "" {
		return "", errors.New("You must provide " + envVar + " in environment variables")
	}

	return value, nil
}

func DisconnectOfDB(MongoClient *mongo.Client) error {
	err := MongoClient.Disconnect(context.TODO())
	if err != nil {
		return err
	}

	return nil
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func CheckPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(password), []byte(hashedPassword))
}
