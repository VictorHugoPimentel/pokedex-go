package main

import (
	"Pokedex/internal/pokeapi"
	"fmt"
	"log"
)

func main() {
	pokeapiClient := pokeapi.NewClient()
	resp, err := pokeapiClient.ListLocationAreas()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
	//startRepl()
}