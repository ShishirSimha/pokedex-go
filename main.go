package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ShishirSimha/pokedex-go/internal/command" // change to your module name
)

func main() {
	command.InitCommands()

	config := &command.Config{}

	fmt.Println("Welcome to the Pokedex! Type 'help' for commands.")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}

		cleanedInput, err := cleanInput(scanner.Text())
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}

		cmdName := cleanedInput[0]
		args := cleanedInput[1:] // remaining words are arguments

		cmd, exists := command.Commands[cmdName]
		if !exists {
			fmt.Println("Unknown command. Type 'help' for available commands.")
			continue
		}

		err = cmd.Callback(config, args)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
}

func cleanInput(text string) ([]string, error) {
	//trim space
	cleanedText := strings.TrimSpace(text)

	if cleanedText == "" {
		return nil, fmt.Errorf("invalid input")
	}

	cleanedText = strings.ToLower(text)
	words := strings.Fields(cleanedText)
	return words, nil
}
