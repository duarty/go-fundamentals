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

	body, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))
	req.Body.Close()
}
