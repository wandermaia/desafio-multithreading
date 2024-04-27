package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	log.Println("Servidor iniciado!")
	http.ListenAndServe(":8000", router)
}

/*

go mod init github.com/wandermaia/desafio-multithreading
go mod tidy

*/
