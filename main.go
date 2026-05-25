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

		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			continue
		}

		cmdName := strings.ToLower(strings.Fields(input)[0])

		cmd, exists := command.Commands[cmdName]
		if !exists {
			fmt.Println("Unknown command. Type 'help' for available commands.")
			continue
		}

		err := cmd.Callback(config)
		if err != nil {
			if err.Error() == "exit" {
				return
			}
			fmt.Printf("Error: %v\n", err)
		}
	}
}
