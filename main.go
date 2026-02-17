package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//Scanner to scan input
	scanner := bufio.NewScanner(os.Stdin)

	//Infinite for loop for the REPL
	for {
		fmt.Print("Pokedex > ")

		//Scan the input now
		scanner.Scan()

		//Register the commands
		commands := RegisterCommands()

		//get the input command and split.
		input := cleanInput(scanner.Text())

		//Check if the first word is a command
		if command, exists := commands[input[0]]; exists {
			//call the command
			command.callback()
		} else {
			printUnknown(input[0])
		}

	}

}

// printUnkown informs the user about invalid commands
func printUnknown(text string) {
	fmt.Println(text, ": command not found")
}
