package main

import (
	"net/http"

	"github.com/wandermaia/desafio-multithreading/internal/infra/webserver/handlers"
)

func main() {
	http.HandleFunc("/", handlers.BuscaCepHandler)
	http.ListenAndServe(":8000", nil)
}

/*

go mod init github.com/wandermaia/desafio-multithreading
go mod tidy


Valida o parâmetro CEP que vem na query da URL.
http://localhost:8000/?cep=32450000


*/
