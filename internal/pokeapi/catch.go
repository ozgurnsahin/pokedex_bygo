package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand/v2"
	"net/http"
	"time"

	"github.com/ozgurnsahin/pokedex_bygo/internal/pokecache"
)

// Explore Locations -
func (c *Client) Catch(pokemon *string) (Pokemon, error) {
	cache := pokecache.NewCache(15 * time.Minute)
	url := baseURL + "/pokemon"

	if pokemon != nil {
		url = baseURL + "/pokemon/" + *pokemon
	}
	random := rand.IntN(2)

	fmt.Printf("Throwing a Pokeball at %v... \n", *pokemon)

	if random == 1 {

		if elem, ok := cache.Entry[url]; !ok {

			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				return Pokemon{}, err
			}

			resp, err := c.httpClient.Do(req)
			if err != nil {
				return Pokemon{}, err
			}
			defer resp.Body.Close()

			dat, err := io.ReadAll(resp.Body)
			if err != nil {
				return Pokemon{}, err
			}

			cache.Add(url,dat)

			var pokemonDetail Pokemon
			err = json.Unmarshal(dat, &pokemonDetail)
			if err != nil {
				return Pokemon{}, err
			}

			fmt.Printf("%v was caught! \n", *pokemon)
			fmt.Println("You may now inspect it with the inspect command.")
			caughtPokemon[*pokemon] = pokemonDetail

			return pokemonDetail, nil

		} else {

			var pokemonDetail Pokemon
			err := json.Unmarshal(elem.Val, &pokemonDetail)
			if err != nil {
				return Pokemon{}, err
			}

			fmt.Printf("%v was caught! \n", *pokemon)
			fmt.Println("You may now inspect it with the inspect command.")
			caughtPokemon[*pokemon] = pokemonDetail

			return pokemonDetail, nil
		}

	}
	
	fmt.Printf("%v escaped! \n", *pokemon)

	return Pokemon{}, nil
	

}

