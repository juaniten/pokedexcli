package main

import (
	"fmt"

	"github.com/juaniten/pokedexcli/internal/pokeapi"
)

func commandMap(config *Config) error {
	if config.Next == "" {
		fmt.Println("there are no more results")
		return nil
	}
	la, err := pokeapi.GetLocations(config.Next)
	if err != nil {
		return err
	}

	pokeapi.PrintLocations(la)

	config.Previous = la.Previous
	config.Next = la.Next

	return nil
}

func commandMapb(config *Config) error {
	if config.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	la, err := pokeapi.GetLocations(config.Previous)
	if err != nil {
		return err
	}

	pokeapi.PrintLocations(la)

	config.Previous = la.Previous
	config.Next = la.Next
	return nil
}
