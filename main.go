package main

import (
	"log"

	"github.com/brunoarruda04/golang-clean-architecture/api"
)

func main() {
	if err := api.NewService().Start(); err != nil {
		log.Fatal(err)
	}
}
