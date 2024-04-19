package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type CEPData struct {
	CEP         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	UF          string `json:"uf"`
	IBGE        string `json:"ibge"`
	GIA         string `json:"gia"`
	DDD         string `json:"ddd"`
	SIAFI       string `json:"siafi"`
}

const URL string = "http://viacep.com.br/ws/cep/json"

func main() {
	for _, cepnumber := range os.Args[1:] {
		prefix := URL[:24]
		sufix := URL[27:]
		println(prefix + cepnumber + sufix)
		req, err := http.Get(prefix + cepnumber + sufix)
		if err != nil {
			panic(err)
		}
		defer req.Body.Close()
		res, err := io.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}

		var cep CEPData
		err = json.Unmarshal(res, &cep)
		if err != nil {
			panic(err)
		}
		fmt.Println(cep.Logradouro)

	}

}
