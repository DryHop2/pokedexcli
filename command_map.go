package main

import (
	"fmt"

	"github.com/DryHop2/pokedexcli/internal/pokeapi"
)

func commandMap(cfg *Config, args []string) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if cfg.NextLocationAreaURL != nil {
		url = *cfg.NextLocationAreaURL
	}

	// resp, err := http.Get(url)
	// if err != nil {
	// 	return fmt.Errorf("failed to fetch location areas: %w", err)
	// }
	// defer resp.Body.Close()

	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	return fmt.Errorf("failed to read response body: %w", err)
	// }

	// var locationData pokeapi.LocationAreaResponse
	// err = json.Unmarshal(body, &locationData)
	// if err != nil {
	// 	return fmt.Errorf("failed to parse JSON: %w", err)
	// }

	// for _, area := range locationData.Results {
	// 	fmt.Println(area.Name)
	// }

	data, err := pokeapi.GetLocationAreas(cfg.Cache, url)
	if err != nil {
		return err
	}

	cfg.NextLocationAreaURL = data.Next
	cfg.PreviousLocationAreaURL = data.Previous

	for _, area := range data.Results {
		fmt.Println(area.Name)
	}

	return nil
}

func commandMapb(cfg *Config, args []string) error {
	if cfg.PreviousLocationAreaURL == nil {
		fmt.Println("You're at the beginning of the map — no previous locations.")
		return nil
	}

	url := *cfg.PreviousLocationAreaURL

	// resp, err := http.Get(url)
	// if err != nil {
	// 	return fmt.Errorf("failed to fetch location areas: %w", err)
	// }
	// defer resp.Body.Close()

	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	return fmt.Errorf("failed to read response body: %w", err)
	// }

	// var locationData pokeapi.LocationAreaResponse
	// err = json.Unmarshal(body, &locationData)
	// if err != nil {
	// 	return fmt.Errorf("failed to parse JSON: %w", err)
	// }

	// for _, area := range locationData.Results {
	// 	fmt.Println(area.Name)
	// }

	data, err := pokeapi.GetLocationAreas(cfg.Cache, url)
	if err != nil {
		return err
	}

	cfg.NextLocationAreaURL = data.Next
	cfg.PreviousLocationAreaURL = data.Previous

	for _, area := range data.Results {
		fmt.Println(area.Name)
	}

	return nil
}
