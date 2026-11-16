package main

import (
	"time"

	"Pokedox/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{
		pokeapiClient: pokeClient,
		caughtPokemon: make(map[string]pokeapi.Pokemon),
		commands:      getCommands(),
	}
	startRepl(cfg)
}
