package handlers

import (
	"context"
	"net/http"

	"github.com/WilliamKSilva/musical-comrade/application/usecases"
	"github.com/WilliamKSilva/musical-comrade/presentation/utils"
)

type CreateGameHandler struct {
	AddGameUseCase usecases.AddGameUseCase
}

func (h CreateGameHandler) CreateGameHandler(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var addGameData usecases.AddGameData

		fields := []string{"genre", "status"}
		message, err := utils.FieldValidator(w, &addGameData, fields)

		if err != nil {
			panic("Error")
		}

		if message != "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(message))
			return
		}

		createdGame, err := h.AddGameUseCase.Add(ctx, &addGameData)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		encodeError := utils.ResponseJSON(w, &createdGame)

		if encodeError != nil {
			panic("Error")
		}
	}
}
