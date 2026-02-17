package main

import (
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	words := strings.Fields(lowered)
	return words
}

//cli command structure
type cliCommand struct {
	name string
	description string
	callback func() error
}

//callback for command exit
func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

//callback for command help
func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: ")

	for command, prop := range RegisterCommands() {
		fmt.Printf("%s: %s\n", command, prop.description)
	}
	return nil
}

func RegisterCommands() map[string]cliCommand {
	commands := map[string]cliCommand {
		"exit" : {
			name : "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"help" : {
			name : "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
	}
	return commands
}


