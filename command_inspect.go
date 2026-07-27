package main

import (
	"fmt"
)

func callbackInspect(cfg *config, args ...string) error {
	pokemon, err := cfg.pokemonService.InspectPokemon(args[0])
	if err != nil {
		return err
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

