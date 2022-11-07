package factories

import (
	"github.com/WilliamKSilva/musical-comrade/application/repositories"
	"github.com/WilliamKSilva/musical-comrade/application/usecases"
	"github.com/WilliamKSilva/musical-comrade/presentation"
	"gorm.io/gorm"
)

func MakeSignUpHandler(db *gorm.DB) presentation.SignUpHandler {
	userRepository := repositories.UserRepositoryDb{Db: db}
	addUserUseCase := usecases.AddUserUseCase{UserRepository: userRepository}
	signUpHandler := presentation.SignUpHandler{AddUserUseCase: addUserUseCase}

	return signUpHandler
}

func MakeCreateGameHandler(db *gorm.DB) presentation.CreateGameHandler {
	gameRepository := repositories.GameRepositoryDb{Db: db}
	addGameUseCase := usecases.AddGameUseCase{GameRepository: gameRepository}
	createGameHandler := presentation.CreateGameHandler{AddGameUseCase: addGameUseCase}

	return createGameHandler
}
