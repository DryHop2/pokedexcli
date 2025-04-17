package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMap(cfg *Config) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if cfg.NextLocationAreaURL != nil {
		url = *cfg.NextLocationAreaURL
	}

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to fetch location areas: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	var locationData LocationAreaResponse
	err = json.Unmarshal(body, &locationData)
	if err != nil {
		return fmt.Errorf("failed to parse JSON: %w", err)
	}

	for _, area := range locationData.Results {
		fmt.Println(area.Name)
	}

	cfg.NextLocationAreaURL = locationData.Next
	cfg.PreviousLocationAreaURL = locationData.Previous

	return nil
}

func commandMapb(cfg *Config) error {
	if cfg.PreviousLocationAreaURL == nil {
		fmt.Println("You're at the beginning of the map — no previous locations.")
		return nil
	}

	url := *cfg.PreviousLocationAreaURL

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to fetch location areas: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	var locationData LocationAreaResponse
	err = json.Unmarshal(body, &locationData)
	if err != nil {
		return fmt.Errorf("failed to parse JSON: %w", err)
	}

	for _, area := range locationData.Results {
		fmt.Println(area.Name)
	}

	// Update pagination state again
	cfg.NextLocationAreaURL = locationData.Next
	cfg.PreviousLocationAreaURL = locationData.Previous

	return nil
}
