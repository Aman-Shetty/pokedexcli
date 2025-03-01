package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no location area provided")
	}
	locationAreaName := args[0]

	locationArea, err := cfg.pokeapiClient.GetLocationArea(locationAreaName)
	if err != nil {
		return nil
	}
	fmt.Printf("Pokemon in %s:\n", locationAreaName)

	for _, pokemon := range locationArea.PokemonEncounters{
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}
	return nil
}
