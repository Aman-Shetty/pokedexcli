package main

import (
	"time"

	pokeapi "github.com/Aman-Shetty/pokedexcli/internal/api"
)

type config struct {
	pokeapiClient pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
}

func main() {
	cfg := config {
		pokeapiClient: pokeapi.NewClient(time.Hour),
	}
	startrepl(&cfg)
}

