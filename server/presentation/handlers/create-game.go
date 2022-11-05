package handlers

import (
	"net/http"

	"github.com/WilliamKSilva/musical-comrade/application/usecases"
	"github.com/WilliamKSilva/musical-comrade/presentation/utils"
)

type CreateGameHandler struct {
	usecase usecases.AddGameUseCase
}

func (h CreateGameHandler) CreateGameHandler(w http.ResponseWriter, req *http.Request) {
	var addGameData usecases.AddGameData

	fields := []string{"genre", "user_id"}
	errorMessage, err := utils.FieldValidator(w, &addGameData, fields)

	if errorMessage != "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(errorMessage))
		return
	}

	if err != nil {
		panic(err)
	}

	jsonParseError := utils.RequestJSON(req, &addGameData)

	if jsonParseError != nil {
		panic(jsonParseError)
	}

	createdGame, err := h.usecase.Add(&addGameData)

	if err != nil {
		panic(err)
	}

}
