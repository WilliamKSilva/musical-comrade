package factories

import (
	"github.com/WilliamKSilva/musical-comrade/application/repositories"
	"github.com/WilliamKSilva/musical-comrade/application/usecases"
	"github.com/WilliamKSilva/musical-comrade/presentation"
	"gorm.io/gorm"
)

func MakeSignUpHandler(db *gorm.DB) presentation.SignUpHandler {
	accountRepository := repositories.AccountRepositoryDb{Db: db}
	addAccountUseCase := usecases.AddAccountUseCase{AccountRepository: accountRepository}
	signUpHandler := presentation.SignUpHandler{AddAccountUseCase: addAccountUseCase}

	return signUpHandler
}
