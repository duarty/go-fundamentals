package main

import (
	"io"
	"net/http"
)

func main() {
	req, err := http.Get("https://google.com.br")
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()
	body, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	defer println("1", string(body))
	println("2", string(body))

}
