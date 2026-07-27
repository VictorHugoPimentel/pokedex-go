package main

import (
	"errors"
)

func callbackCatch(cfg *config, args ...string) error {

    if len(args) != 1 {
        return errors.New("No Pokemon name provided.")
    }

    return cfg.pokemonService.CatchPokemon(args[0])
}

