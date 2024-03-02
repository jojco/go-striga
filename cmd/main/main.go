package main

import (
	"log"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func main() {
	log.Println("Starting go-stringa ...")

	mux := http.NewServeMux()
	mux.HandleFunc("/", helloWorld)

	connStr := ":8080"
	server := &http.Server{
		Addr:    connStr,
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
	log.Println("Server listening on: ", connStr)
}
