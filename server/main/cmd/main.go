package main

import (
	"fmt"
	"net/http"

	"github.com/WilliamKSilva/musical-comrade/infra/db"
	"github.com/WilliamKSilva/musical-comrade/main/factories"
	"github.com/gorilla/mux"
)

func main() {
	connectionDb, err := db.Connect()

	if err != nil {
		fmt.Println("Error connecting to database")
	}

	router := mux.NewRouter()

	makeSignUpHandler := factories.MakeSignUpHandler(connectionDb)

	router.HandleFunc("/user", makeSignUpHandler.SignUpHandler)

	fmt.Println("Running!")

	defer fmt.Println(http.ListenAndServe(":3000", router))
}
