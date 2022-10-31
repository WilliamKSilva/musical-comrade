package usecases

import (
	"github.com/WilliamKSilva/musical-comrade/application/repositories"
	"github.com/WilliamKSilva/musical-comrade/domain"
	"github.com/google/uuid"
)

type AddAccountData struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
}

type AddAccountUseCase struct {
	AccountRepository repositories.AccountRepository
}

func (repo AddAccountUseCase) Add(account *AddAccountData) (*domain.User, error) {
	uuid := uuid.New().String()

	user := domain.User{
		Id:       uuid,
		Name:     account.Name,
		Email:    account.Name,
		Password: account.Password,
		Nickname: account.Password,
	}

	createdUser, err := repo.AccountRepository.Add(&user)

	if err != nil {
		return nil, err
	}

	return createdUser, nil
}
