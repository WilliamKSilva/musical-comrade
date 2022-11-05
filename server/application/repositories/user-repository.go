package repositories

import (
	"github.com/WilliamKSilva/musical-comrade/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	Add(user *domain.User) (*domain.User, error)
}

type UserRepositoryDb struct {
	Db *gorm.DB
}

func (database UserRepositoryDb) Add(user *domain.User) (*domain.User, error) {

	err := database.Db.Create(user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}
