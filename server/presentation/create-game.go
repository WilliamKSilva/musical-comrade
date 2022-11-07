package presentation

import (
	"net/http"

	"github.com/WilliamKSilva/musical-comrade/application/usecases"
	"github.com/WilliamKSilva/musical-comrade/presentation/utils"
)

type CreateGameHandler struct {
	AddGameUseCase usecases.AddGameUseCase
}

func (h CreateGameHandler) CreateGameHandler(w http.ResponseWriter, req *http.Request) {
	var gameData usecases.AddGameData

	fields := []string{"user_id, status"}
	message, err := utils.FieldValidator(w, &gameData, fields)

	if err != nil {
		panic(err)
	}

	if message != "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(message))
		return
	}

	parseRequestError := utils.RequestJSON(req, &gameData)

	if parseRequestError != nil {
		panic(parseRequestError)
	}

	createdGame, err := h.AddGameUseCase.Add(&gameData)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	utils.ResponseJSON(w, &createdGame)
}
