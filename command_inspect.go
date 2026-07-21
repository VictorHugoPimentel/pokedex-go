package main

import (
	"errors"
	"fmt"
)

func callbackInspect(cfg *config, args ...string) error {
	if len(args) != 1{
		return errors.New("No Pokemon name provided.")
	}

	pokemonName := args[0]
	pokemon, ok := cfg.caughtPokemons[pokemonName]
	if !ok{
		return fmt.Errorf("You haven't caught %s yet.", pokemonName)
	}
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	for _, types := range pokemon.Types {
		fmt.Printf("Type: %s\n", types.Type.Name)
	}
	fmt.Printf("ID: %d\n",pokemon.ID)
	fmt.Println("Stats:")
	for _, stats := range pokemon.Stats {
		fmt.Printf("%s: %d\n", stats.Stat.Name, stats.BaseStat)
	}
	return nil
}

