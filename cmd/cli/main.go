package main

import (
	"Pokedex/internal/pokeapi"
	"Pokedex/internal/service"
	"time"
)

type config struct {
	pokemonService *service.PokemonService
	locationService *service.LocationService
}

func main() {
	client := pokeapi.NewClient(time.Hour)

	cfg := config{
		pokemonService: service.NewPokemonService(client),
		locationService: service.NewLocationService(client),
	}
	startRepl(&cfg)
}