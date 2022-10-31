package repositories

import (
	"github.com/WilliamKSilva/musical-comrade/domain"
	"gorm.io/gorm"
)

type AccountRepository interface {
	Add(user *domain.User) (*domain.User, error)
}

type AccountRepositoryDb struct {
	Db *gorm.DB
}

func (repo AccountRepositoryDb) Add(user *domain.User) (*domain.User, error) {

	err := repo.Db.Create(user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}
