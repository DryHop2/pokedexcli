package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/DryHop2/pokedexcli/internal/pokecache"
)

type LocationAreaResponse struct {
	Count    int                `json:"count"`
	Next     *string            `json:"next"`
	Previous *string            `json:"previous"`
	Results  []LocationAreaInfo `json:"results"`
}

type LocationAreaInfo struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func GetLocationAreas(cache *pokecache.Cache, url string) (LocationAreaResponse, error) {
	var data LocationAreaResponse

	if val, ok := cache.Get(url); ok {
		err := json.Unmarshal(val, &data)
		if err != nil {
			return data, fmt.Errorf("error decoding cached response: %w", err)
		}
		return data, nil
	}

	resp, err := http.Get(url)
	if err != nil {
		return data, fmt.Errorf("failed to fetch location areas: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return data, fmt.Errorf("failed to read response body: %w", err)
	}

	cache.Add(url, body)

	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return data, nil
}
