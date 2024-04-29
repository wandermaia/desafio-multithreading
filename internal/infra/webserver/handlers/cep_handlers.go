package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

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
	Site        string
}

type BrasilApiCep struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
	Site         string
}

func BuscaCepHandler(w http.ResponseWriter, r *http.Request) {

	cepParam := chi.URLParam(r, "cep")
	cep, err := BuscaCepBrasilApi(cepParam)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Erro ao buscar o cep: %s", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(cep)

}

func BuscaCepViaCep(cep string) (*ViaCEP, error) {

	url := "http://viacep.com.br/ws/" + cep + "/json/"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var dadosCep ViaCEP
	err = json.Unmarshal(body, &dadosCep)
	if err != nil {
		return nil, err
	}

	dadosCep.Site = url

	return &dadosCep, nil

}

func BuscaCepBrasilApi(cep string) (*BrasilApiCep, error) {

	url := "https://brasilapi.com.br/api/cep/v1/" + cep
	resp, err := http.Get(url)
	if err != nil {
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
	dadosCep.Site = url

	return &dadosCep, nil
}
