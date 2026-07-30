package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("What do you want to do? Type 'help' for a list of commands: ")
		scanner.Scan()
		text := scanner.Text()
		cleaned := cleanInput(text)
		if len(cleaned) == 0{
			continue
		}
		commandName := cleaned[0]
		args := []string{}
		if len(cleaned) > 1{
			args = cleaned[1:]
		}

		availableCommands := getCommands()
		command, ok := availableCommands[commandName]
		if !ok{
			fmt.Println("Invalid command. Please try again.")
			continue
		}
		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}
	}
}

type cliCommand struct {
	name string
	description string
	callback func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name: "help",
			description: "Show this help message",
			callback: callbackHelp,
		},
		"exit": {
			name: "exit",
			description: "Exit the REPL",
			callback: callbackExit,
		},
		"map": {
			name: "map",
			description: "List all location areas",
			callback: callbackMap,
		},
		"mapb": {
			name: "mapb",
			description: "List previous location areas",
			callback: callbackMapb,
		},
		"explore": {
			name: "explore {location_area_name}",
			description: "Explore Pokemons in a location area",
			callback: callbackExplore,
		},
		"catch": {
			name: "catch {pokemon_name}",
			description: "Catch a Pokemon",
			callback: callbackCatch,
		},
		"inspect": {
			name: "inspect {pokemon_name}",
			description: "Inspect a Pokemon",
			callback: callbackInspect,
		},
		"pokedex": {
			name: "pokedex",
			description: "Shows the details of a caught Pokemon",
			callback: callbackPokedex,
		},
	}
}

func cleanInput(str string)[]string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}