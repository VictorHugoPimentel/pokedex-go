package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func callbackCatch(cfg *config, args ...string) error {
	if len(args) != 1{
		return errors.New("No Pokemon name provided.")
	}

	pokemonName := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}
	const threshold = 50
	randnum := rand.Intn(pokemon.BaseExperience)
	fmt.Println(pokemon.BaseExperience, randnum, threshold)
	if randnum > threshold{
		return fmt.Errorf("Failed to catch %s", pokemon.Name)
	}
	cfg.caughtPokemons[pokemon.Name] = pokemon
	fmt.Printf("Caught %s!\n", pokemon.Name)
	return nil
}

