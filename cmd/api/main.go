package main

import (
	"Pokedex/internal/pokeapi"
	"Pokedex/internal/service"
	"fmt"
	"net/http"
	"time"
)

type Server struct {
	pokemonService *service.PokemonService
	locationService *service.LocationService
}

func main() {
	client := pokeapi.NewClient(time.Hour)

	server := &Server{
		pokemonService: service.NewPokemonService(client),
		locationService: service.NewLocationService(client),
	}
	
	http.HandleFunc("/pokemon/", server.pokemonHandler)
	http.HandleFunc("/ping", server.pingHandler)

	fmt.Println("Server running on :8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}