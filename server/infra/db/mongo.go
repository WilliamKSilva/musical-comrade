package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDb(ctx context.Context) *mongo.Client {

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://docker:admin@localhost:27020"))

	if err != nil {
		panic(err)
	}

	return client
}
