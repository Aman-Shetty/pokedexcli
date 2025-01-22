package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}
	pokemonName := args[0]

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return nil
	}

	const threshold = 50
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	randomNumber := rand.Intn(pokemon.BaseExperience)
	if randomNumber > threshold {
		return fmt.Errorf("%s escaped", pokemonName)
	}

	cfg.caughtPokemon[pokemonName] = pokemon
	fmt.Printf("%s was caught!\n", pokemonName)
	return nil
}
