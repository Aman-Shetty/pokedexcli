package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config) error {
	response, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)
	if err != nil {
		return nil
	}
	fmt.Println("Location Areas: ")

	for _, areas := range response.Results{
		fmt.Printf(" - %s\n", areas.Name)
	}
	cfg.nextLocationAreaURL = response.Next
	cfg.prevLocationAreaURL = response.Previous
	return nil
}

func commandMapb(cfg *config) error {
	if cfg.prevLocationAreaURL == nil {
		return errors.New("you're on the first page")
	}
	response, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationAreaURL)
	if err != nil {
		return err
	}
	fmt.Println("Location Areas: ")

	for _, areas := range response.Results{
		fmt.Printf(" - %s\n", areas.Name)
	}
	cfg.nextLocationAreaURL = response.Next
	cfg.prevLocationAreaURL = response.Previous
	return nil
}