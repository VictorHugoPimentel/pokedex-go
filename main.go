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
	caughtPokemons map[string]pokeapi.Pokemon
}

func main() {
	client := pokeapi.NewClient(time.Hour)

	cfg := config{
		pokemonService: service.NewPokemonService(client),
	}
	startRepl(&cfg)
}