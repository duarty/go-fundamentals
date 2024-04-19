package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Account struct {
	Number  int
	Balance float64
}

// json annotation  json -> struct
type Account2 struct {
	Number  int     `json:"n"`
	Balance float64 `json:"s"`
	Debit   float64 `json:"-"` //ignore field
}

func main() {
	acc := Account{
		Number:  1,
		Balance: 0.10,
	}

	res, err := json.Marshal(acc)
	if err != nil {
		panic(err)
	}
	println(string(res))

	err = json.NewEncoder(os.Stdout).Encode(acc)
	if err != nil {
		panic(err)
	}

	jsonFile := []byte(`{"number": 2, "balance": 0.20}`)
	var accX Account
	json.Unmarshal(jsonFile, &accX)
	if err != nil {
		panic(err)
	}
	fmt.Println(accX)

	jsonFile2 := []byte(`{"n": 2, "s": 0.20}`)
	var accX2 Account2
	json.Unmarshal(jsonFile2, &accX2)
	if err != nil {
		panic(err)
	}
	fmt.Println(accX2)
}
