package usecases

import (
	"github.com/WilliamKSilva/musical-comrade/application/repositories"
	"github.com/WilliamKSilva/musical-comrade/domain"
	"github.com/google/uuid"
)

type AddGameData struct {
	UserId string                `json:"user_id"`
	Genre  domain.GameMusicGenre `json:"genre"`
}

type AddGameUseCase struct {
	GameRepository repositories.GameRepository
}

func (repo AddGameUseCase) Add(game *AddGameData) (*domain.Game, error) {

	gameData := domain.Game{
		Id:     uuid.New().String(),
		Genre:  game.Genre,
		UserId: game.UserId,
	}

	createdGame, err := repo.GameRepository.Add(&gameData)

	if err != nil {
		return nil, err
	}

	return createdGame, nil
}
