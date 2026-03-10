package main

import()

type Config struct {
	previous *string
	next *string
}
//https://pokeapi.co/api/v2/location-area/1
//
func main() {
	cfg := &Config{
		previous: nil,
		next: nil,	
	}
	startREPL(cfg)
}

