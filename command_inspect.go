package main

import (
	"errors"
	"fmt"
	"strings"
)

func commandInspect(cfg *Config, args []string) error {
	if len(args) < 1 {
		return errors.New("you must provide a Pokemon name")
	}

	pokemonName := strings.ToLower(args[0])
	pokemon, ok := cfg.Pokedex[pokemonName]
	if !ok {
		fmt.Println("you have not caught that pokemon yet!")
		return nil
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)

	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("   -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("   - %s\n", t.Type.Name)
	}

	return nil
}
