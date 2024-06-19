package database

import (
	"context"
	"fmt"
	"log"
	"os"

	//"github.com/AlejandroWaiz/server/model"

	"github.com/AlejandroWaiz/server/model"
	"google.golang.org/api/iterator"
)

func (db *Database) GetAllPokemons() ([]model.FrontendPokedexPokemon, error) {

	ctx := context.Background()

	collection := db.dbClient.Collection(os.Getenv("pokeapi_pokemons_collection"))

	iter := collection.Documents(ctx)

	//var pokemons []model.PokeapiPokemon
	var pokemons []model.FrontendPokedexPokemon
	var pokemon model.FrontendPokedexPokemon
	var errors []error

	defer iter.Stop() // add this line to ensure resources cleaned up
	var pok int

	for {

		doc, err := iter.Next()

		pok++

		if err == iterator.Done {
			break
		}

		if pok == 3 {
			break
		}

		if err != nil {

			if len(errors) > 20 {
				log.Println(err)
				errors = append(errors, fmt.Errorf("[Database] Err getting document: %v", err))
				continue
			} else {
				break
			}

		}

		//TODO: decidido cambiar todo el funcionamiento a la funcion "DataAt" para buscar solo la informacion requerida
		//Y no toda por cada documento, asique habrá que agregar individualmente los campos y manejar sus errores

		pokemonData := doc.Data()

		if err != nil {
			errors = append(errors, fmt.Errorf("[Database] Err maping doc data at path: %v", err))
			continue
		}

		pokemon.Name = pokemonData["Name"].(string)
		pokemon.ID = pokemonData["ID"].(int64)

		StatsData, err := doc.DataAt("Stats")
		pokemon.Stats = StatsData

		typesData, err := doc.DataAt("Types")
		pokemon.ElementalTypes = typesData

		pokemons = append(pokemons, pokemon)

	}

	if len(errors) > 0 {
		return pokemons, fmt.Errorf("Err getting all pokemons")
	} else {
		return pokemons, nil
	}

}
