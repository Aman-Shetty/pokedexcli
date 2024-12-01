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
	callback func() error 
}

func startrepl() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("> ")
		reader.Scan()
		
		words := clearInputs(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		command, ok := getCommands()[commandName]

		if ok {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
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