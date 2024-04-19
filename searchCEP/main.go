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

func main() {
	for _, cepnumber := range os.Args[1:] {
		req, err := http.Get("http://viacep.com.br/ws/" + cepnumber + "/json")
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

		file, err := os.Create("cep.txt")
		_, err = file.WriteString(cep.Logradouro)
		//os.Remove("cep.txt")
		fmt.Println(cep.Logradouro)
	}

}
