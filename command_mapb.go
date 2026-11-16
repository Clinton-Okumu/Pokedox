package main

import (
	"fmt"
)

func commandMapb(cfg *config, args ...string) error {
	if cfg.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	resp, err := cfg.pokeapiClient.ListLocations(cfg.Previous)
	if err != nil {
		return err
	}

	cfg.Next = resp.Next
	cfg.Previous = resp.Previous

	for _, loc := range resp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
