package repositories

import (
	"github.com/WilliamKSilva/musical-comrade/domain"
	"gorm.io/gorm"
)

type GameRepository interface {
	Add(game *domain.Game) (*domain.Game, error)
}

type GameRepositoryDb struct {
	Db *gorm.DB
}

func (repo GameRepositoryDb) Add(game *domain.Game) (*domain.Game, error) {
	result := repo.Db.Save(game)

	if result.Error != nil {
		return nil, result.Error
	}

	return game, nil
}
