package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type clicommand struct {
	name	string
	description string
	callback func(*config, ...string) error 
}

func startrepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("> ")
		reader.Scan()
		
		words := clearInputs(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, ok := getCommands()[commandName]

		if !ok {
			fmt.Println("Unknown command")
			continue
			
		} 
		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func getCommands() map[string]clicommand {
	commands := map[string]clicommand {
		"help": {
			name: "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name: "map",
			description: "Lists the location areas",
			callback:    commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Lists the location areas backwards",
			callback:    commandMapb,
		},
		"explore": {
			name: "explore {location_area}",
			description: "Lists the pokemon in a location area",
			callback:    commandExplore,
		},
		"catch": {
			name: "catch {pokemon_name}",
			description: "Attempt to catch a pokemon and add to your pokedex",
			callback:    commandCatch,
		},
		"inspect": {
			name: "inspect {pokemon_name}",
			description: "Gives the details about the pokemon",
			callback:    commandInspect,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
	return commands
}

func clearInputs(text string) []string {
	text = strings.ToLower(text)
	words := strings.Fields(text)
	return words
}