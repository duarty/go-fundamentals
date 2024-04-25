package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8000", nil)
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()
	io.Copy(os.Stdout, req.Body)
}
