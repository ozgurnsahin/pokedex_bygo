package main

import (
	"fmt"
)

func commandPokedex(cfg *config) error {

	pokemon_list := cfg.pokeapiClient.Pokedex()

	fmt.Print("Your Pokedex: \n")
	for _,poke := range pokemon_list {

		fmt.Printf(" -%s \n", poke.Name)

	}

	return nil
}