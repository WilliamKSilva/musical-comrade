package repositories

import (
	"github.com/WilliamKSilva/musical-comrade/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	Add(user *domain.User) (*domain.User, error)
}

type UserRepositoryDb struct {
	Db *mongo.Database
}

func (database UserRepositoryDb) Add(user *domain.User) (*domain.User, error) {

	err := database.Db.Collection(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}
