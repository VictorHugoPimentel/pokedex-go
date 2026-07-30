package service

import (
	"Pokedex/internal/pokeapi"
	"fmt"
	"math/rand"
	"sort"
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

func (s *PokemonService) InspectPokemon(name string) (pokeapi.Pokemon, error) {
	pokemon, ok := s.caught[name]
	if !ok {
		return pokeapi.Pokemon{}, fmt.Errorf("You haven't caught %s yet.", name)
	}
	return pokemon, nil
}

func (s*PokemonService) ListCaughtPokemons() []pokeapi.Pokemon {
	pokemons := make([]pokeapi.Pokemon, 0, len(s.caught))
	for _, pokemon := range s.caught {
		pokemons = append(pokemons, pokemon)
	}
	sort.Slice(pokemons, func(i, j int) bool {
		return pokemons[i].ID < pokemons[j].ID
	})
	return pokemons
}

func (s *PokemonService) GetPokemon(name string) (pokeapi.Pokemon, error) {
	return s.client.GetPokemon(name)
}