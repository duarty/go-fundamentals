package main

import "net/http"

func main() {
	mux := http.NewServeMux()

	//mux.HandleFunc("/", mux1Handle)
	mux.Handle("/mux1", blog{title: "mux1"})
	http.ListenAndServe(":8080", mux)

	mux2 := http.NewServeMux()
	mux2.HandleFunc("/mux2", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("mux2")) })
	http.ListenAndServe(":8081", mux2)
}

type blog struct {
	title string
}

// func mux1Handle(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("mux1"))
// }

func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.title))
}
