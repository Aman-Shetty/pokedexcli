package main

import "fmt"

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Available commands are: ")
	for _, cmd := range getCommands() {
		fmt.Printf("\t%s: %s \n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}