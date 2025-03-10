//"github.com/jojco/go-striga/pkg/pkg2"
//"github.com/jojco/go-striga/pkg/pkg3"
//"github.com/jojco/go-striga/pkg/pkg1"

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
	w.Write([]byte("Welcome to Striga page!"))

}

func helloWorld2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("2"))

}

func main() {
	log.Println("Starting go-stringa ...")

	pkg3.teploty()
	pkg2.udajescd30()
	pkg1.rele()

	mux := http.NewServeMux()
	mux.HandleFunc("/", helloWorld)
	mux.HandleFunc("/2", helloWorld2)

	// Handle interrupt signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	// Start HTTP server test zmeny wwefwefwe
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
