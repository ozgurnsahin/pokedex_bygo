package main

import (
	"fmt"
)

func commandExplore(cfg *config) error {
	explore, err := cfg.pokeapiClient.Explore(cfg.areaName)
	if err != nil {
		return err
	}

	for _, loc := range explore.PokemonEncounters {
		fmt.Println(loc.Pokemon.Name)
	}

	return nil
}
