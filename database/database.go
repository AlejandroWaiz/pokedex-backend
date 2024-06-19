package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"github.com/AlejandroWaiz/server/model"
	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
	"google.golang.org/api/option"
)

type Database struct {
	app      *firebase.App
	dbClient *firestore.Client
}

type DatabaseImplementation interface {
	GetAllPokemons() ([]model.FrontendPokedexPokemon, error)
	createDatabase() []error
}

func New() (DatabaseImplementation, error) {

	firestoreApp, err := firebase.NewApp(context.Background(), nil, option.WithCredentialsFile(os.Getenv("service_account_path")))

	if err != nil {
		err = fmt.Errorf("[Database] Error initializing firebase: %v\n", err)
		return nil, err
	}

	firestoreClient, err := firestoreApp.Firestore(context.Background())

	if err != nil {
		err = fmt.Errorf("[Database] Error initializing firestore: %v\n", err)
		return nil, err
	}

	db := Database{app: firestoreApp, dbClient: firestoreClient}

	// dbErrors := db.createDatabase()

	// if len(dbErrors) > 0 {

	// 	for i, err := range dbErrors {

	// 		log.Printf("Err n° %v: %v", i, err)

	// 	}

	// 	return nil, errors.New("ERROR CREATING DATABASE")

	// }

	return &db, nil

}

func (d *Database) createDatabase() []error {

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
			log.Printf("[Pokeapi] Err getting %v data: %v", r.Name, err)
			errors = append(errors, err)
		}

		_, err := d.dbClient.Collection(os.Getenv("pokeapi_pokemons_collection")).Doc(strconv.Itoa(i)).Set(context.Background(), pokemon)

		if err != nil {
			log.Printf("[Firestore] Err saving %v data: %v", r.Name, err)
			errors = append(errors, err)
		}

	}

	return errors

}
