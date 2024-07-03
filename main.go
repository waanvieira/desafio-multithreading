package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Cep struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}

type CepViaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	ch1 := make(chan Cep)
	ch2 := make(chan CepViaCep)

	go BuscaBrasilApi(ch1)
	go BuscaViaCep(ch2)

	select {
	case msg1 := <-ch1:
		fmt.Println("received ", msg1)
		fmt.Printf("https://brasilapi.com.br/api/cep/v1/01310930")
	case msg2 := <-ch2:
		fmt.Println("received", msg2)
		fmt.Printf("http://viacep.com.br/ws/01310930/json")
	case <-time.After(time.Second):
		fmt.Println("timeout")
	}
}

func BuscaBrasilApi(ch chan Cep) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", "https://brasilapi.com.br/api/cep/v1/01310930", nil)
	if err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	data, error := io.ReadAll(res.Body)
	if error != nil {
		log.Fatal(error)
	}

	var apiRes Cep
	error = json.Unmarshal([]byte(data), &apiRes)
	if error != nil {
		panic(error)
	}

	ch <- apiRes
}

func BuscaViaCep(ch chan CepViaCep) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", "http://viacep.com.br/ws/01310930/json/", nil)
	if err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	data, error := io.ReadAll(res.Body)
	if error != nil {
		log.Fatal(error)
	}

	var apiRes CepViaCep
	error = json.Unmarshal([]byte(data), &apiRes)
	if error != nil {
		panic(error)
	}

	ch <- apiRes
}
