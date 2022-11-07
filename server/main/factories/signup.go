package factories

import (
	"github.com/WilliamKSilva/musical-comrade/application/repositories"
	"github.com/WilliamKSilva/musical-comrade/application/usecases"
	"github.com/WilliamKSilva/musical-comrade/presentation/handlers"

	"go.mongodb.org/mongo-driver/mongo"
)

func MakeSignUpHandler(db *mongo.Database) handlers.SignUpHandler {
	userRepository := repositories.UserRepositoryDb{Db: db}
	addUserUseCase := usecases.AddUserUseCase{UserRepository: userRepository}
	signUpHandler := handlers.SignUpHandler{AddUserUseCase: addUserUseCase}

	return signUpHandler
}
