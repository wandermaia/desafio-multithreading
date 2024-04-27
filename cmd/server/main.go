package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/wandermaia/desafio-multithreading/internal/infra/webserver/handlers"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	//router.Get("/{cep}", han)

}

/*

go mod init github.com/wandermaia/desafio-multithreading
go mod tidy

*/
