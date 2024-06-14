package database

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

type Database struct {
	app      *firebase.App
	dbClient *firestore.Client
}

type DatabaseImplementation interface {
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

	return &Database{app: firestoreApp, dbClient: firestoreClient}, nil

}
