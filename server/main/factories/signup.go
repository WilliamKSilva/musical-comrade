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

func MakeCreateGameHandler(db *mongo.Database) presentation.CreateGameHandler {
	gameRepository := repositories.GameRepositoryDb{Db: db}
	addGameUseCase := usecases.AddGameUseCase{GameRepository: gameRepository}
	createGameHandler := presentation.CreateGameHandler{AddGameUseCase: addGameUseCase}

	return createGameHandler
}
