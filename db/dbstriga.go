package dbstriga

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func DbStriga() {
	//názov s veľkým začiatočným písmenom dovoľuje volanie funkcie
	//aj z iných balíkov, a teda funkcia je verejná

	// Otvorenie alebo vytvorenie databázy
	db, err := sql.Open("sqlite3", "./striga.db")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Vytvorenie tabuľky teplôt, ak neexistuje
	_, err = db.Exec(`
                CREATE TABLE IF NOT EXISTS teploty (
                        id INTEGER PRIMARY KEY AUTOINCREMENT,
                        sensorid   TEXT// Device ID
						hodnota REAL,
                        cas DATETIME DEFAULT CURRENT_TIMESTAMP
                )
        `)

	if err != nil {
		log.Fatal(err)
	}

	// Vytvorenie tabuľky vlhkosti, ak neexistuje
	_, err = db.Exec(`
                CREATE TABLE IF NOT EXISTS vlhkost (
                        id INTEGER PRIMARY KEY AUTOINCREMENT,
                        hodnota REAL,
                        cas DATETIME DEFAULT CURRENT_TIMESTAMP
                )
        `)

	if err != nil {
		log.Fatal(err)
	}

	// Vloženie 10 hodnôt teploty
	for i := 1; i <= 10; i++ {
		teplomer := fmt.Sprintf("28-%d", i)
		teplota := 20.0 + rand.Float64()*10.0 // Generovanie náhodnej teploty medzi 20 a 30 stupňami

		_, err = db.Exec("INSERT INTO teploty (teplomer, teplota) VALUES (?, ?)", teplomer, teplota)
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(1 * time.Second) // odstup 1 sekunda medzi meraniami.
	}

	fmt.Println("Dáta boli úspešne zapísané do databázy.")
}
