package database

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go/v4"
)

type Database struct {
	db *firebase.App
}

type DatabaseImplementation interface {
}

func New() (DatabaseImplementation, error) {

	db, err := firebase.NewApp(context.Background(), nil)

	if err != nil {
		err = fmt.Errorf("[Database] Error initializing: %v\n", err)
		return nil, err
	}

	return &Database{db: db}, nil

}
