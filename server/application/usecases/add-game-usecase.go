package usecases

import (
	"github.com/WilliamKSilva/musical-comrade/application/repositories"
	"github.com/WilliamKSilva/musical-comrade/domain"
	"github.com/google/uuid"
)

type AddGameData struct {
	Status domain.GameStatus `json:"status"`
	UserId string            `json:"user_id"`
}

type AddGameUseCase struct {
	GameRepository repositories.GameRepository
}

func (repo AddGameUseCase) Add(gameData *AddGameData) (*domain.Game, error) {

	game := domain.Game{
		Id:     uuid.New().String(),
		Status: gameData.Status,
		UserId: gameData.UserId,
	}

	createdGame, err := repo.GameRepository.Add(&game)

	if err != nil {
		return nil, err
	}

	return createdGame, nil

}
