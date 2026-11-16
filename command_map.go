package main

import (
	"fmt"
)

func commandMap(cfg *config, args ...string) error {
	resp, err := cfg.pokeapiClient.ListLocations(cfg.Next)
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
