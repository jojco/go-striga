//"github.com/jojco/go-striga/pkg/pkg1"
//"github.com/jojco/go-striga/pkg/pkg2"
//"github.com/jojco/go-striga/pkg/pkg3"

package main

import (
	"log"

	"github.com/jojco/go-striga/pkg/pkg1"
	"github.com/jojco/go-striga/pkg/pkg2"
	"github.com/jojco/go-striga/pkg/pkg3"
	webserver "github.com/jojco/go-striga/web"
)

func main() {
	log.Println("Starting go-stringa ...")

	pkg3.Meranieteploty()
	pkg2.Udajezscd30()
	pkg1.OvladanieRele()
	webserver.Webserverstriga()

}
