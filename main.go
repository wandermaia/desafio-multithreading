package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

// Constraint criada para facilitar a função de formatação do json para exibição
type DadosCep interface {
	*ViaCEP | *BrasilApiCep
}

// Struct para coleta de dados a partir do site do ViaCEP
type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

// Struct para coleta de dados a partir do site do BrasilApi
type BrasilApiCep struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}

func main() {

	// Verifica se o CEP foi passado por parâmetro
	if len(os.Args) < 2 {
		fmt.Println("O CEP deve ser pasado por parâmetro.")
		os.Exit(1)
	}

	cepParam := os.Args[1]
	fmt.Printf("CEP Informado: %s\n\n", cepParam)

	// Validando se o CEP informado é válido.
	if validarParametro(cepParam) {
		fmt.Printf("O CEP %s é válido.\n\n", cepParam)
	} else {
		fmt.Printf("O CEP %s é inválido. Deve ter exatamente 8 caracteres e ser composto apenas por números inteiros.\n\n", cepParam)
		os.Exit(1)
	}

	// Canais que serão utilizados para realizar a passagem dos valores entre as threads
	channel1 := make(chan *BrasilApiCep)
	channel2 := make(chan *ViaCEP)

	// Função anônima que realiza a busca do CEP no site BrasilAPI e salva o resultado no channel01
	go func() {
		cep, err := BuscaCepBrasilApi(cepParam)

		if err != nil {
			fmt.Printf("Erro ao consultar cep BrasilAPI: %s", err)
			panic(err)
		}
		//time.Sleep(time.Second * 2)
		channel1 <- cep

	}()

	// Função anônima que realiza a busca do CEP no site ViaCep e salva o resultado no channel02
	go func() {
		cep2, err := BuscaCepViaCep(cepParam)
		if err != nil {
			fmt.Printf("Erro ao consultar BuscaCepViaCep: %s", err)
		}
		time.Sleep(time.Second * 2)
		channel2 <- cep2
	}()

	select {
	case dadosCep := <-channel1: // BuscaCepBrasilApi
		fmt.Printf("Dados recebidos do site https://brasilapi.com.br: \n\n%s\n\n", FormataJson(dadosCep))

	case dadosCep := <-channel2: // BuscaCepViaCep
		fmt.Printf("Dados recebidos do site https://viacep.com.br: \n\n%s\n\n", FormataJson(dadosCep))

	case <-time.After(time.Second):
		println("Timeout! As APIs demoraram mais do que 1 segundo para responder!")
	}

}

// Função para formatação do json que será exibido na tela
func FormataJson[T DadosCep](dados T) string {
	jsonData, err := json.MarshalIndent(dados, "", "  ")
	if err != nil {
		fmt.Println("Erro ao formatar o JSON:", err)
		panic(err)
	}
	return string(jsonData)
}

// Função que realiza busca no site https://viacep.com.br
func BuscaCepViaCep(cep string) (*ViaCEP, error) {

	url := "http://viacep.com.br/ws/" + cep + "/json/"
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	log.Println(resp.StatusCode)

	var dadosCep ViaCEP
	err = json.Unmarshal(body, &dadosCep)
	if err != nil {
		return nil, err
	}
	return &dadosCep, nil

}

// Função que realiza busca no site https://brasilapi.com.br
func BuscaCepBrasilApi(cep string) (*BrasilApiCep, error) {

	url := "https://brasilapi.com.br/api/cep/v1/" + cep
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var dadosCep BrasilApiCep
	err = json.Unmarshal(body, &dadosCep)
	if err != nil {
		return nil, err
	}
	return &dadosCep, nil
}

// Função que valida o CEP informado por parâmetro
func validarParametro(parametro string) bool {
	// Verifica se o parâmetro tem exatamente 8 caracteres
	if len(parametro) != 8 {
		return false
	}

	// Verifica se todos os caracteres são números inteiros
	_, err := strconv.Atoi(parametro)
	return err == nil
}
