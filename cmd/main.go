//"github.com/jojco/go-striga/pkg/pkg1"
//"github.com/jojco/go-striga/pkg/pkg2"
//"github.com/jojco/go-striga/pkg/pkg3"

package main

import (
	//"log"

	"fmt"

	dbstriga "github.com/jojco/go-striga/db"
	"github.com/jojco/go-striga/pkg/pkg1"
	"github.com/jojco/go-striga/pkg/pkg2"
	"github.com/jojco/go-striga/pkg/pkg3"
	webserver "github.com/jojco/go-striga/web"
)

func main() {

	for {

		// Tu môžeš pridať kód, ktorý sa má vykonávať v slučke
		fmt.Println("Slučka beží...")
		pkg3.Meranieteploty()
		pkg2.Udajezscd30()
		pkg1.OvladanieRele()
		webserver.Webserverstriga()
		dbstriga.DbStriga()
	}
}
