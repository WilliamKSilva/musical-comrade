package repositories

import (
	"context"

	"github.com/WilliamKSilva/musical-comrade/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	Add(ctx context.Context, user *domain.User) (*domain.User, error)
}

type UserRepositoryDb struct {
	Db *mongo.Database
}

func (repo UserRepositoryDb) Add(ctx context.Context, user *domain.User) (*domain.User, error) {

	repo.Db.CreateCollection(ctx, "users")

	inserted, err := repo.Db.Collection("users").InsertOne(ctx, user)

	if err != nil {
		panic(err)
	}

	createdUser := repo.Db.Collection("users").FindOne(ctx, bson.M{"_id": inserted.InsertedID})

	if createdUser.Err() != nil {
		return nil, createdUser.Err()
	}

	parsedUser := domain.User{}

	createdUser.Decode(&parsedUser)

	if err != nil {
		return nil, err
	}

	return &parsedUser, nil
}
