package main

import (
	"time"

	"github.com/DryHop2/pokedexcli/internal/pokecache"
)

type Config struct {
	NextLocationAreaURL     *string
	PreviousLocationAreaURL *string
	Cache                   *pokecache.Cache
}

func NewConfig() *Config {
	return &Config{
		Cache: pokecache.NewCache(5 * time.Second),
	}
}
