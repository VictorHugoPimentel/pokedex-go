package main

import (
	"Pokedex/internal/pokeapi"
	"Pokedex/internal/service"
	"time"
)

type config struct {
	pokeapiClient pokeapi.Client
	pokemonService *service.PokemonService
	nextLocationAreaURL *string
	prevLocationAreaURL *string
	locationService *service.LocationService
	caughtPokemons map[string]pokeapi.Pokemon
}

func main() {
	client := pokeapi.NewClient(time.Hour)

	cfg := config{
		pokemonService: service.NewPokemonService(client),
		locationService: service.NewLocationService(client),
	}
	startRepl(&cfg)
}