package repositories

import (
	"github.com/WilliamKSilva/musical-comrade/domain"
	"gorm.io/gorm"
)

type GameRepository interface {
	Add(*domain.Game) (*domain.Game, error)
}

type GameRepositoryDb struct {
	Db *gorm.DB
}

func (database GameRepositoryDb) Add(game *domain.Game) (*domain.Game, error) {

	err := database.Db.Save(game).Error

	if err != nil {
		return nil, err
	}

	return game, nil
}
