package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/wandermaia/desafio-multithreading/internal/infra/webserver/handlers"
)

func main() {

	//cepHandler := handlers.GetCepHandler()
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	//router.Get("/{cep}",)

}

/*

go mod init github.com/wandermaia/desafio-multithreading
go mod tidy

*/
