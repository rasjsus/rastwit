package main

import (
	"log"

	"github.com/rasjsus/rastwit/db"
	"github.com/rasjsus/rastwit/handlers"
)

func main() {
	if db.CheckConnection() == 0 {
		log.Fatal("No hay conneccion a la base de datos")
		return
	}
	handlers.Handlers()
}
