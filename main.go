package main

import (
	"log"

	"github.com/AlejandroWaiz/server/database"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load(".env")

	db, err := database.New()

	if err != nil {
		log.Printf("Err: %v", err)
	}

	errors := db.CreateDatabase()

	if len(errors) > 0 {

		for i, err := range errors {
			log.Printf("err n° %v: %v", i, err)
		}

	}

}
