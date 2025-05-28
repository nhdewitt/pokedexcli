package main

import (
	"time"

	"github.com/nhdewitt/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := &config{
		pokeapiClient: pokeClient,
		pokedex:       Pokedex{caught: make(map[string]Pokemon)},
	}

	startRepl(cfg)
}
