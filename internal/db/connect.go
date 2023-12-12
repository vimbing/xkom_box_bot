package db

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	err      error
	DbClient *mongo.Client
)

func Connect() error {
	uri := os.Getenv("MONGODB_URI")

	DbClient, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		return err
	}

	fmt.Println("Successfully connected to db!")

	return nil
}

func Disconnect() error {
	return DbClient.Disconnect(context.TODO())
}
