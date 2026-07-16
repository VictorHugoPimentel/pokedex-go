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

		availableCommands := getCommands()
		command, ok := availableCommands[commandName]
		if !ok{
			fmt.Println("Invalid command. Please try again.")
			continue
		}
		err := command.callback(cfg)
		if err != nil {
			fmt.Println(err)
		}
	}
}

type cliCommand struct {
	name string
	description string
	callback func(*config) error
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
	}
}

func cleanInput(str string)[]string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}