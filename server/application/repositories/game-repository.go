package repositories

import (
	"context"

	"github.com/WilliamKSilva/musical-comrade/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type GameRepository interface {
	Add(ctx context.Context, game *domain.Game) (*domain.Game, error)
}

type GameRepositoryDb struct {
	Db *mongo.Database
}

func (repo GameRepositoryDb) Add(ctx context.Context, game *domain.Game) (*domain.Game, error) {

	repo.Db.CreateCollection(ctx, "games")

	inserted, err := repo.Db.Collection("games").InsertOne(ctx, game)

	if err != nil {
		panic(err)
	}

	createdGame := repo.Db.Collection("games").FindOne(ctx, bson.M{"_id": inserted.InsertedID})

	if createdGame.Err() != nil {
		return nil, err
	}

	parsedGame := domain.Game{}
	createdGame.Decode(&parsedGame)

	return &parsedGame, nil
}
