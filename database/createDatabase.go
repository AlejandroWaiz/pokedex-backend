package database

import (
	"context"
	"log"
	"os"
	"strconv"

	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
)

func (d *Database) CreateDatabase() []error {

	var errors []error

	resources, err := pokeapi.Resource("pokemon", 0, 1300)

	if err != nil {
		log.Printf("[Database] Err getting pokeapi resources: %v", err)
		errors = append(errors, err)
		return errors
	}

	var pokemon structs.Pokemon

	for i, r := range resources.Results {

		pokemon, err = pokeapi.Pokemon(r.Name)

		if err != nil {
			log.Printf("Err getting %v data: %v", r.Name, err)
			errors = append(errors, err)
		}

		_, err := d.dbClient.Collection(os.Getenv("pokeapi_pokemons_collection")).Doc(strconv.Itoa(i)).Set(context.Background(), pokemon)

		if err != nil {
			log.Printf("Err saving %v data: %v", r.Name, err)
			errors = append(errors, err)
		}

	}

	return errors

}
