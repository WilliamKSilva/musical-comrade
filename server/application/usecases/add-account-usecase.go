package usecases

import (
	"github.com/WilliamKSilva/musical-comrade/application/repositories"
	"github.com/WilliamKSilva/musical-comrade/domain"
)

type AddAccountUseCase struct {
	AccountRepository repositories.AccountRepository
}

func (repo AddAccountUseCase) Add(user *domain.User) (*domain.User, error) {
	createdUser, err := repo.Add(user)

	if err != nil {
		return nil, err
	}

	return createdUser, nil
}
