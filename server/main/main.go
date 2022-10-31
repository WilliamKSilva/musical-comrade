package main

import (
	"fmt"
	"net/http"

	"github.com/WilliamKSilva/musical-comrade/infra/db"
	"github.com/gorilla/mux"
)

func main() {
	connectionDb, err := db.Connect()

	if err != nil {
		fmt.Println("Error connecting to database")
	}

	router := mux.NewRouter()

	router.HandleFunc("/user", MakeSignUpHandler(connectionDb))

	fmt.Println(http.ListenAndServe(":3000", router))
}
