package repositories

import (
	"fmt"

	"github.com/WilliamKSilva/musical-comrade/domain"
	"gorm.io/gorm"
)

type GameRepository interface {
	Add(*domain.Game) (*domain.Game, error)
}

type GameRepositoryDb struct {
	database *gorm.DB
}

func (repo GameRepositoryDb) GameRepository(game *domain.Game) (*domain.Game, error) {
	result := repo.database.Save(game)

	fmt.Println(game)

	if result.Error != nil {
		return nil, result.Error
	}

	return game, nil
}
