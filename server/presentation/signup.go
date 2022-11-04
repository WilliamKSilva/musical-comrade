package presentation

import (
	"net/http"

	"github.com/WilliamKSilva/musical-comrade/application/usecases"
	utils "github.com/WilliamKSilva/musical-comrade/presentation/utils"
)

type SignUpHandler struct {
	AddAccountUseCase usecases.AddAccountUseCase
}

func (h *SignUpHandler) SignUpHandler(w http.ResponseWriter, req *http.Request) {
	var addAccountData usecases.AddAccountData
	err := utils.RequestJSON(req, &addAccountData)

	if err != nil {
		panic("Error")
	}

	fields := []string{"name", "email", "password", "nickname"}
	errorMessage, err := utils.FieldValidator(w, &addAccountData, fields)

	if errorMessage != "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(errorMessage))
		return
	}

	if err != nil {
		panic("Error")
	}

	createdUser, err := h.AddAccountUseCase.Add(&addAccountData)

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
