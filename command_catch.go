package main

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/DryHop2/pokedexcli/internal/pokeapi"
)

func commandCatch(cfg *Config, args []string) error {
	if len(args) < 1 {
		return errors.New("you must provide a Pokemon name")
	}

	pokemonName := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	pokemon, err := pokeapi.GetPokemon(cfg.Cache, pokemonName)
	if err != nil {
		return fmt.Errorf("failed to fetch Pokemon data: %w", err)
	}

	// rand.Seed(time.Now().UnixNano())
	// catchThreshold := pokemon.BaseExperience

	difficulty := pokemon.BaseExperience
	catchChance := 100 - (difficulty / 2)

	if catchChance < 5 {
		catchChance = 5
	}

	chance := rand.Intn(100)

	if chance < catchChance {
		fmt.Printf("%s was caught!\n", pokemonName)
		cfg.Pokedex[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s escaped !\n", pokemonName)
	}

	// if chance < catchThreshold {
	// 	fmt.Printf("%s escaped!\n", pokemonName)
	// 	return nil
	// }

	// fmt.Printf("%s was caught!\n", pokemonName)
	// cfg.Pokedex[pokemon.Name] = pokemon

	return nil
}
