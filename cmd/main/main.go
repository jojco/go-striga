package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func main() {
	log.Println("Starting go-stringa ...")

	mux := http.NewServeMux()
	mux.HandleFunc("/", helloWorld)

	// Handle interrupt signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	// Start HTTP server test zmenz wwefwefwe
	server := &http.Server{Addr: ":8080", Handler: mux}

	log.Default().Println("Server is listening on port 8080...")

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Error starting server: %s\n", err)
		}
	}()
	<-interrupt

	log.Default().Println("Shutting down server...")
	if err := server.Shutdown(nil); err != nil {
		log.Fatal(fmt.Printf("Error shutting down server: %s\n", err))
	}
	log.Default().Println("Server stopped.")
}
