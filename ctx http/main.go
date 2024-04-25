package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}
func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Inicou")
	defer log.Println("Finalizada")
	select {
	case <-time.After(5 * time.Second):
		// Imprime no servidor
		log.Println("Req processada")
		// client
		w.Write([]byte("processada"))
	case <-ctx.Done():
		log.Println("Req fnalizada pelo cliente")
	}

}
