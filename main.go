package main

import (
	"log"

	"github.com/brunoarruda04/golang-clean-architecture/api"
	"github.com/brunoarruda04/golang-clean-architecture/infra/config"
)

func main() {
	var err error
	err = config.StartConfig()
	FatalError(err)

	err = api.NewService().Start()
	FatalError(err)
}

func FatalError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
