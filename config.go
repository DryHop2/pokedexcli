package main

import (
	"time"

	"github.com/DryHop2/pokedexcli/internal/pokeapi"
	"github.com/DryHop2/pokedexcli/internal/pokecache"
)

type Config struct {
	NextLocationAreaURL     *string
	PreviousLocationAreaURL *string
	Cache                   *pokecache.Cache
	Pokedex                 map[string]pokeapi.Pokemon
}

func NewConfig() *Config {
	return &Config{
		Cache:   pokecache.NewCache(5 * time.Second),
		Pokedex: make(map[string]pokeapi.Pokemon),
	}
}
