package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}
	pokemonName := args[0]

	pokemon, ok := cfg.caughtPokemon[pokemonName]

	if !ok {
		return errors.New("you haven't caught this pokemon yet")
	}
	
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("BaseExperience: %d\n", pokemon.BaseExperience)
	fmt.Println("Stats:")
	for _, stats := range pokemon.Stats {
		fmt.Printf("\t-%s: %v\n", stats.Stat.Name, stats.BaseStat)
	}
	fmt.Println("Types:")
	for _, types := range pokemon.Types {
		fmt.Printf("\t-%s\n", types.Type.Name)
	}
	// fmt.Println("Moves:")
	// for _, move := range pokemon.Moves {
	// 	fmt.Printf("\t-%s\n", move.Move.Name)
	// }
	return nil
}
