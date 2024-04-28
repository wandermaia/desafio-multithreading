package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/wandermaia/desafio-multithreading/internal/infra/webserver/handlers"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Get("/{cep}", handlers.BuscaCepHandler)
	//router.Get("/", handlers.BuscaCepHandler)

	log.Println("Servidor iniciado!")
	http.ListenAndServe(":8000", router)
}

/*

go mod init github.com/wandermaia/desafio-multithreading
go mod tidy


Valida o parâmetro CEP que vem na query da URL.
http://localhost:8000/?cep=32450000


*/
