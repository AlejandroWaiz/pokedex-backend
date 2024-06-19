package database

import (
	"context"
	"encoding/json"
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
	var results []model.FrontendPokedexPokemon
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

		// var pokemon model.PokeapiPokemon
		// if err := doc.DataTo(&pokemon); err != nil {

		// 	if len(errors) > 20 {
		// 		log.Println(err)
		// 		errors = append(errors, fmt.Errorf("[Database] Err maping doc data: %v", err))
		// 		continue
		// 	} else {
		// 		break
		// 	}

		// }

		pokemonData := doc.Data()

		if err != nil {
			errors = append(errors, fmt.Errorf("[Database] Err maping doc data at path: %v", err))
			continue
		}

		pokemon.Name = pokemonData["Name"].(string)
		pokemon.ID = pokemonData["ID"].(int64)
		//Stats := pokemonData["Stats"].([]model.PokeapiStat)

		StatsData, err := doc.DataAt("Abilities")

		if err != nil {
			errors = append(errors, err)
		}

		var Stats []model.PokeapiStat

		err = json.Unmarshal([]byte(fmt.Sprintf("%v", StatsData)), &Stats)

		if err != nil {
			errors = append(errors, err)
		}

		if len(Stats) == 0 {
			errors = append(errors, fmt.Errorf("Null info from firestore"))
		}

		for _, s := range Stats {
			var stat model.PokemonStat
			stat.Name = s.Stat.Name
			stat.BaseStat = s.BaseStat
			pokemon.Stats = append(pokemon.Stats, stat)
		}

		results = append(results, pokemon)

		//pokemons = append(pokemons, pokemon)

	}

	if len(errors) > 0 {
		return results, fmt.Errorf("Err getting all pokemons")
	} else {
		return results, nil
	}

}
