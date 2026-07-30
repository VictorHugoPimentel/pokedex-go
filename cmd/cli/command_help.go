package main

import "fmt"

var commandOrder = []string{
	"catch",
	"inspect",
	"pokedex",
	"map",
	"mapb",
	"explore",
	"help",
	"exit",
}

func callbackHelp(cfg *config, args ...string) error {
	fmt.Println("Welcome to TOG's POKEDEX!!")
	fmt.Println("Here are the available commands:")

	availableCommands := getCommands()

	for _, commandName := range commandOrder {
		command, ok := availableCommands[commandName]
		if ! ok {
			continue
		}
		fmt.Printf("%s - %s\n", command.name, command.description)
	}

	fmt.Println()
	return nil
}