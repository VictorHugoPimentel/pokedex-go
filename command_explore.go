package main

import (
	"fmt"
)

func callbackExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("Please provide a valid location area name.")
	}
	locationArea, err := cfg.locationService.ExploreLocationArea(args[0])
	if err != nil {
		return err
	}
	fmt.Printf("Pokemon in %s:\n", locationArea.Name)
	for _, pokemon := range locationArea.PokemonEncounters {
		fmt.Println("-", pokemon.Pokemon.Name)
	}
	return nil
}

