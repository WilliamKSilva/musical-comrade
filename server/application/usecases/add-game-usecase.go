package usecases

import (
	"context"

	"github.com/WilliamKSilva/musical-comrade/application/repositories"
	"github.com/WilliamKSilva/musical-comrade/domain"
)

type AddGameData struct {
	Status string `json:"status"`
	Genre  string `json:"genre"`
}

type AddGameUseCase struct {
	GameRepository repositories.GameRepository
}

func (repo AddGameUseCase) Add(ctx context.Context, game *AddGameData) (*domain.Game, error) {

	gameData := domain.Game{
		Genre:  domain.Genre(game.Genre),
		Status: domain.Status(game.Status),
	}

	createdGame, err := repo.GameRepository.Add(ctx, &gameData)

	if err != nil {
		return nil, err
	}

	return createdGame, nil
}
