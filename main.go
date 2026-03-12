package main

import(
	"pokedexcli/internal/pokeapi"
	"time"
)

type Config struct {
	previous *string
	next *string
	client pokeapi.Client
}
//https://pokeapi.co/api/v2/location-area/1
//
func main() {
	cfg := &Config{
		previous: nil,
		next: nil,
		client: pokeapi.NewClient(5*time.Second, 5*time.Minute),
	}
	startREPL(cfg)
}

