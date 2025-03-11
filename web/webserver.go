package webserver

import (
	"fmt"
	"net/http"
)

// Handler pre hlavnú stránku
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Ahoj, vitajte na Raspberry Pi web serveri!")
}

// Funkcia na spustenie servera
func StartServer(port string) error {
	// Definujeme router
	http.HandleFunc("/", homePage)

	// Spustíme server na zadanom porte
	fmt.Println("Server beží na http://localhost" + port)
	return http.ListenAndServe(port, nil)
}



func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to Striga page!"))

}

func helloWorld2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("2"))

}

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

// Spracovanie signálov na ukončenie
quit := make(chan os.Signal, 1)
signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
<-quit
log.Println("Server is shutting down...")

// Vytvorenie kontextu s časovým limitom
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

log.Default().Println("Shutting down server...")
// Vypnutie servera s kontextom
if err := server.Shutdown(ctx); err != nil {
	log.Fatalf("Server forced to shutdown: %v", err)
}

log.Default().Println("Server stopped.")