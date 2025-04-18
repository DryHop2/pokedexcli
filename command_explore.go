package main

import (
	"errors"
	"fmt"

	"github.com/DryHop2/pokedexcli/internal/pokeapi"
)

func commandExplore(cfg *Config, args []string) error {
	if len(args) < 1 {
		return errors.New("you must provide a location area name")
	}

	areaName := args[0]
	url := "https://pokeapi.co/api/v2/location-area/" + areaName

	data, err := pokeapi.GetLocationArea(cfg.Cache, url)
	if err != nil {
		return err
	}

	fmt.Println("Exploring", areaName, "...")
	fmt.Println("Found Pokemon:")

	for _, p := range data.PokemonEncounters {
		fmt.Println("-", p.Pokemon.Name)
	}

	return nil
}
