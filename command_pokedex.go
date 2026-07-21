package main

import (
	"fmt"
)

func callbackPokedex(cfg *config, args ...string) error {
	fmt.Println("Here is your Pokedex!!!")
	if len(cfg.caughtPokemons) == 0 {
		fmt.Println("It's Empty! You haven't caught any Pokemon yet.")
		return nil
	}
	for _, pokemon := range cfg.caughtPokemons {
		fmt.Printf("Name: %s - ID: %d\n", pokemon.Name, pokemon.ID)
	}
	return nil
}

