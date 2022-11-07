package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/WilliamKSilva/musical-comrade/infra/db"
	"github.com/WilliamKSilva/musical-comrade/main/factories"
	"github.com/gorilla/mux"
)

func main() {
	ctx := context.Background()
	mongoClient := db.ConnectDb(ctx)

	database := mongoClient.Database("musical-comrade-db")

	defer mongoClient.Disconnect(ctx)

	router := mux.NewRouter()

	makeSignUpHandler := factories.MakeSignUpHandler(database)
	makeCreateGameHandler := factories.MakeCreateGameHandler(database)

	router.HandleFunc("/users", makeSignUpHandler.SignUpHandler)
	router.HandleFunc("/games", makeCreateGameHandler.CreateGameHandler)

	defer fmt.Println(http.ListenAndServe(":3000", router))
}
