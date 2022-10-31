package usecases

import (
	"github.com/WilliamKSilva/musical-comrade/application/repositories"
	"github.com/WilliamKSilva/musical-comrade/domain"
	"github.com/google/uuid"
)

type AddAccountUseCase struct {
	AccountRepository repositories.AccountRepository
}

func (repo AddAccountUseCase) Add(user *domain.User) (*domain.User, error) {
	uuid := uuid.New().String()

	user.Id = uuid

	createdUser, err := repo.AccountRepository.Add(user)

	if err != nil {
		return nil, err
	}

	return createdUser, nil
}
