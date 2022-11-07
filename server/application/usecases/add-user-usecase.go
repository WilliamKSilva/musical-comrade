package usecases

import (
	"context"

	"github.com/WilliamKSilva/musical-comrade/application/repositories"
	"github.com/WilliamKSilva/musical-comrade/domain"
)

type AddUserData struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
}

type AddUserUseCase struct {
	UserRepository repositories.UserRepository
}

func (repo AddUserUseCase) Add(ctx context.Context, account *AddUserData) (*domain.User, error) {

	user := domain.User{
		Name:     account.Name,
		Email:    account.Email,
		Password: account.Password,
		Nickname: account.Nickname,
	}

	createdUser, err := repo.UserRepository.Add(ctx, &user)

	if err != nil {
		return nil, err
	}

	return createdUser, nil
}
