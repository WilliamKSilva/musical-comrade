package presentation

import (
	"net/http"

	"github.com/WilliamKSilva/musical-comrade/application/usecases"
	utils "github.com/WilliamKSilva/musical-comrade/presentation/utils"
)

type SignUpHandler struct {
	AddUserUseCase usecases.AddUserUseCase
}

func (h *SignUpHandler) SignUpHandler(w http.ResponseWriter, req *http.Request) {
	var addUserData usecases.AddUserData
	err := utils.RequestJSON(req, &addUserData)

	if err != nil {
		panic("Error")
	}

	fields := []string{"name", "email", "password", "nickname"}
	errorMessage, err := utils.FieldValidator(w, &addUserData, fields)

	if errorMessage != "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(errorMessage))
		return
	}

	if err != nil {
		panic("Error")
	}

	createdUser, err := h.AddUserUseCase.Add(&addUserData)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encodeError := utils.ResponseJSON(w, err.Error())

		if encodeError != nil {
			panic("Error")
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	encodeError := utils.ResponseJSON(w, &createdUser)

	if encodeError != nil {
		panic("Error")
	}

}
