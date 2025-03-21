package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ozgurnsahin/pokedex_bygo/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	areaName *string
	pokemonName *string
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		command, exists := getCommands()[commandName]
		if exists {

			if commandName == "explore" {
				if  len(words) > 1 && len(words[1]) > 0  {
					area := words[1]
					cfg.areaName = &area
				} else {
					fmt.Println("Error: Please provide an area name")
					continue
				}
			} else if commandName == "catch" {
				if  len(words) > 1 && len(words[1]) > 0 {
					pokemon := words[1]
					cfg.pokemonName = &pokemon
				} else {
					fmt.Println("Error: Please provide a Pokemon name")
					continue
				}
			} 

			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}

			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)

	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"explore": {
			name:        "explore",
			description: "Explore the given map",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch a Pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a Pokemon you caught",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Inspect your pokedex",
			callback:    commandPokedex,
		},
	}
}
