package main

import (
	"pokedexcli/internal/pokeapi"
	"time"
)

type Config struct {
	previous   *string
	next       *string
	client     pokeapi.Client
	poke_index map[string]pokeapi.Pokemon
}

// https://pokeapi.co/api/v2/location-area/1
func main() {
	cfg := &Config{
		previous:   nil,
		next:       nil,
		client:     pokeapi.NewClient(5*time.Second, 5*time.Minute),
		poke_index: make(map[string]pokeapi.Pokemon),
	}
	startREPL(cfg)
}
