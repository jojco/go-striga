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
