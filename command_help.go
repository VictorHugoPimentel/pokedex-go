package main

import (
	"fmt"
	"sort"
)

func callbackHelp(cfg *config, args ...string) error {
	fmt.Println("Welcome to TOG's POKEDEX!!")
	fmt.Println("Here are the available commands:")
	availableCommands := getCommands()
	keys := make([]string, 0, len(availableCommands))
	for key := range availableCommands {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		command := availableCommands[key]
		fmt.Printf("%s - %s\n", command.name, command.description)
	}
	fmt.Println()
	return nil
}