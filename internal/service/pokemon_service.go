package service

import (
	"Pokedex/internal/pokeapi"
	"fmt"
	"math/rand"
)

type PokemonService struct {
    client pokeapi.Client
    caught map[string]pokeapi.Pokemon
}

func NewPokemonService(client pokeapi.Client) *PokemonService {
    return &PokemonService{
        client: client,
        caught: make(map[string]pokeapi.Pokemon),
    }
}

func(s *PokemonService) CatchPokemon(name string) error {
	pokemon, err := s.client.GetPokemon(name)
	if err != nil {
		return err
	}
	const threshold = 50
	randnum := rand.Intn(pokemon.BaseExperience)
	fmt.Println(pokemon.BaseExperience, randnum, threshold)
	if randnum > threshold{
		return fmt.Errorf("Failed to catch %s", pokemon.Name)
	}
	s.caught[pokemon.Name] = pokemon
	fmt.Printf("Caught %s!\n", pokemon.Name)
	return nil
}
