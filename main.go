package main

import (
	"log"
	"redSocial/bd"
	"redSocial/handlers"
)

func main() {
	if bd.CheckConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Mapping()
}
