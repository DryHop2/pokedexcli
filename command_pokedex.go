package main

import (
	"fmt"
	"sort"
)

func commandPokedex(cfg *Config, args []string) error {
	if len(cfg.Pokedex) == 0 {
		fmt.Println("You haven't caught any Pokemon yet")
		return nil
	}

	fmt.Println("Your Pokemon: ")

	names := make([]string, 0, len(cfg.Pokedex))
	for name := range cfg.Pokedex {
		names = append(names, name)
	}

	sort.Strings(names)

	for _, name := range names {
		fmt.Printf("- %s\n", name)
	}

	return nil
}
