package presentation

import (
	"net/http"

	"github.com/WilliamKSilva/musical-comrade/application/usecases"
)

type SignUpHandler struct {
	AddAccountUseCase usecases.AddAccountUseCase
}

func (h *SignUpHandler) SignUpHandler(w http.ResponseWriter, req *http.Request) {
	var addAccountData usecases.AddAccountData
	err := requestJSON(req, &addAccountData)

	if err != nil {
		panic("Error")
	}

	createdUser, err := h.AddAccountUseCase.Add(&addAccountData)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(nil)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	encodeError := responseJSON(w, &createdUser)

	if encodeError != nil {
		panic("Error")
	}

}
