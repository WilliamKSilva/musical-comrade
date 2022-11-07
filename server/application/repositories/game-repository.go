package repositories

import (
	"github.com/WilliamKSilva/musical-comrade/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type GameRepository interface {
	Add(game *domain.Game) (*domain.Game, error)
}

type GameRepositoryDb struct {
	Db *mongo.Database
}

func (repo GameRepositoryDb) Add(game *domain.Game) (*domain.Game, error) {
	result := repo.Db.Collection(game)

	if result.Error != nil {
		return nil, result.Error
	}

	return game, nil
}
