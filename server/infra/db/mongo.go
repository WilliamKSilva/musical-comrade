package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDb(ctx context.Context) *mongo.Client {

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://musical_comrade_db"))

	if err != nil {
		panic(err)
	}

	return client
}
