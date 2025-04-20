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

type PokemonEncounter struct {
	Pokemon struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"Pokemon"`
}

type LocationArea struct {
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

type Pokemon struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"Height"`
	Weight         int    `json:"Weight"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}

func GetLocationAreas(cache *pokecache.Cache, url string) (LocationAreaResponse, error) {
	// var data LocationAreaResponse

	// if val, ok := cache.Get(url); ok {
	// 	err := json.Unmarshal(val, &data)
	// 	if err != nil {
	// 		return data, fmt.Errorf("error decoding cached response: %w", err)
	// 	}
	// 	return data, nil
	// }

	// resp, err := http.Get(url)
	// if err != nil {
	// 	return data, fmt.Errorf("failed to fetch location areas: %w", err)
	// }
	// defer resp.Body.Close()

	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	return data, fmt.Errorf("failed to read response body: %w", err)
	// }

	// cache.Add(url, body)

	// err = json.Unmarshal(body, &data)
	// if err != nil {
	// 	return data, fmt.Errorf("failed to parse JSON: %w", err)
	// }

	// return data, nil
	return fetchJSON[LocationAreaResponse](cache, url)
}

func GetLocationArea(cache *pokecache.Cache, url string) (LocationArea, error) {
	// cachedData, ok := cache.Get(url)
	// if ok {
	// 	var data LocationArea
	// 	if err := json.Unmarshal(cachedData, &data); err == nil {
	// 		return data, nil
	// 	}
	// }

	// resp, err := http.Get(url)
	// if err != nil {
	// 	return LocationArea{}, err
	// }
	// defer resp.Body.Close()

	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	return LocationArea{}, err
	// }

	// cache.Add(url, body)

	// var result LocationArea
	// if err := json.Unmarshal(body, &result); err != nil {
	// 	return LocationArea{}, err
	// }

	// return result, nil
	return fetchJSON[LocationArea](cache, url)
}

func GetPokemon(cache *pokecache.Cache, name string) (Pokemon, error) {
	url := "https://pokeapi.co/api/v2/pokemon/" + name

	// if data, ok := cache.Get(url); ok {
	// 	var p Pokemon
	// 	if err := json.Unmarshal(data, &p); err == nil {
	// 		return p, nil
	// 	}
	// }

	// resp, err := http.Get(url)
	// if err != nil {
	// 	return Pokemon{}, err
	// }
	// defer resp.Body.Close()

	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	return Pokemon{}, err
	// }

	// cache.Add(url, body)

	// var p Pokemon
	// if err := json.Unmarshal(body, &p); err != nil {
	// 	return Pokemon{}, err
	// }

	// return p, nil
	return fetchJSON[Pokemon](cache, url)
}

func fetchJSON[T any](cache *pokecache.Cache, url string) (T, error) {
	var result T

	if data, ok := cache.Get(url); ok {
		err := json.Unmarshal(data, &result)
		if err == nil {
			return result, nil
		}
	}

	resp, err := http.Get(url)
	if err != nil {
		return result, nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return result, fmt.Errorf("request failed with status %d (%s)", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	cache.Add(url, body)

	err = json.Unmarshal(body, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}
