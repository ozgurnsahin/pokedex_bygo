package main

import (
	"fmt"
)

func commandInspect(cfg *config) error {

	pokemon := cfg.pokeapiClient.Inspect(cfg.pokemonName)
	if len(pokemon.Name) > 0 {
		fmt.Printf("  Name:            %s\n", pokemon.Name)
		fmt.Printf("  Base Experience: %d\n", pokemon.BaseExperience)
		fmt.Printf("  Height:          %d \n", pokemon.Height)
		fmt.Printf("  Weight:          %d \n", pokemon.Weight)
		
		fmt.Println("  Types:")
		if len(pokemon.Types) == 0 {
			fmt.Println("    None")
		} else {
			for _, typeInfo := range pokemon.Types {
				fmt.Printf("    - %s\n", typeInfo.Type.Name)
			}
		}
	} else {
	fmt.Println("you have not caught that pokemon")
	}

	return nil
}