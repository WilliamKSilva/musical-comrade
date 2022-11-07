package factories

import (
	"github.com/WilliamKSilva/musical-comrade/application/repositories"
	"github.com/WilliamKSilva/musical-comrade/application/usecases"
	"github.com/WilliamKSilva/musical-comrade/presentation"
	"go.mongodb.org/mongo-driver/mongo"
)

func MakeSignUpHandler(db *mongo.Database) presentation.SignUpHandler {
	userRepository := repositories.UserRepositoryDb{Db: db}
	addUserUseCase := usecases.AddUserUseCase{UserRepository: userRepository}
	signUpHandler := presentation.SignUpHandler{AddUserUseCase: addUserUseCase}

	return signUpHandler
}
