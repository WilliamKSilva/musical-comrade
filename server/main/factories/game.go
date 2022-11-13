package factories

import (
	"github.com/WilliamKSilva/musical-comrade/application/repositories"
	"github.com/WilliamKSilva/musical-comrade/application/usecases"
	"github.com/WilliamKSilva/musical-comrade/presentation/handlers"
	"go.mongodb.org/mongo-driver/mongo"
)

func MakeCreateGameHandler(db *mongo.Database) handlers.CreateGameHandler {
	gameRepository := repositories.GameRepositoryDb{Db: db}
	addGameUseCase := usecases.AddGameUseCase{GameRepository: gameRepository}
	createGameHandler := handlers.CreateGameHandler{AddGameUseCase: addGameUseCase}

	return createGameHandler
}
