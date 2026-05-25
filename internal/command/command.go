package command

import (
	"fmt"
	"github.com/ShishirSimha/pokedex-go/internal/util"
)

// Config holds the state for pagination
type Config struct {
    Next     *string
    Previous *string
}

type CLICommand struct {
    Name        string
    Description string
    Callback    func(*Config) error
}

// Global commands map
var Commands = make(map[string]CLICommand)

// Initialize all commands
func InitCommands() {
    Commands["help"] = CLICommand{
        Name:        "help",
        Description: "Displays a help message",
        Callback:    commandHelp,
    }
    Commands["exit"] = CLICommand{
        Name:        "exit",
        Description: "Exit the Pokedex",
        Callback:    commandExit,
    }
    Commands["quit"] = CLICommand {
	    Name : "quit",
	    Description: "Exit the Pokedex",
	    Callback: commandExit,
    }
    Commands["map"] = CLICommand{
        Name:        "map",
        Description: "Display the next 20 location areas",
        Callback:    commandMap,
    }
    Commands["mapb"] = CLICommand {
	    Name: "mapb",
	    Description: "Display the previous 20 location areas",
	    Callback: commandMapBack,
    }
}

// Helper to register more commands later
func RegisterCommand(name, desc string, cb func(*Config) error) {
    Commands[name] = CLICommand{Name: name, Description: desc, Callback: cb}
}

//Definition of the command handlers

//Command handler for Help
func commandHelp(cfg *Config) error {
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Println("Available commands:")
	for _, cmd := range Commands {
		fmt.Printf("  %s - %s\n", cmd.Name, cmd.Description)
	}
	fmt.Println()
	return nil
}

// commandExit
func commandExit(cfg *Config) error {
	fmt.Println("Exiting Pokedex... Goodbye!")
	return fmt.Errorf("exit") // special error to signal exit
}

// commandMap - The main map command
func commandMap(cfg *Config) error {
	url := "https://pokeapi.co/api/v2/location-area?limit=20"
	if cfg.Next != nil && *cfg.Next != "" {
		url = *cfg.Next
	}

	resp, err := util.FetchLocationAreas(url)
	if err != nil {
		return fmt.Errorf("failed to fetch locations: %w", err)
	}

	// Print locations
	for _, area := range resp.Results {
		fmt.Println(area.Name)
	}

	// Update config for next pagination
	cfg.Next = resp.Next
	cfg.Previous = resp.Previous

	return nil
}

//commandMapb - backward Map command
func commandMapBack(cfg *Config) error {
	//Validate if we are in the first page
	if cfg.Previous == nil || *cfg.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	//Now call the API
	resp, err := util.FetchLocationAreas(*cfg.Previous)
	if err != nil {
		return fmt.Errorf("failed to fetch location: %w", err)
	}

	//Print Locations
	for _, area := range resp.Results {
		fmt.Println(area.Name)
	}

	//Update the pointers for pagination
	cfg.Next = resp.Next
	cfg.Previous = resp.Previous

	return nil
}
