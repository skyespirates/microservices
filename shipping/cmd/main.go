package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /shipping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("initialize shipping service"))
	})

	mux.HandleFunc("POST /shipping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("shipping new order"))
	})

	server := http.Server{
		Addr:    ":3003",
		Handler: mux,
	}

	log.Println("server running on port 3003")
	log.Println(server.ListenAndServe())
}
