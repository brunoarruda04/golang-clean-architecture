package main

import (
	"context"
	"log"

	"github.com/brunoarruda04/golang-clean-architecture/api"
	"github.com/brunoarruda04/golang-clean-architecture/infra/config"
	"github.com/brunoarruda04/golang-clean-architecture/infra/database"
	"github.com/brunoarruda04/golang-clean-architecture/infra/database/mongo"
)

func main() {
	var err error
	ctx := context.TODO()

	err = config.StartConfig()
	FatalError(err)

	db := GetDatabase(ctx)

	err = api.NewService(db).Start()
	FatalError(err)
}

func FatalError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func GetDatabase(ctx context.Context) *database.Database {
	client, err := mongo.GetConnection(ctx)
	FatalError(err)

	return database.NewDatabase(client)
}
