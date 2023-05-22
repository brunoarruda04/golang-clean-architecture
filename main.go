package main

import "github.com/brunoarruda04/golang-clean-architecture/api"

func main() {
	service := api.NewService()
	service.Start()
}
