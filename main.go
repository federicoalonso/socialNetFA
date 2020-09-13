package main

import (
	"log"

	"github.com/federicoalonso/socialnetfa/config"
	"github.com/federicoalonso/socialnetfa/routes"
)

func main() {
	if !config.CheckConnection() {
		log.Fatal("Sin conexion a BD")
		return
	}
	routes.Routes()
}
