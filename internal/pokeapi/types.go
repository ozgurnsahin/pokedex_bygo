package pokeapi

var caughtPokemon = make(map[string]Pokemon)

// RespShallowLocations -
type RespShallowLocations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type RespExploreMap struct {
	ID                int    `json:"id"`
    Name              string `json:"name"`
    PokemonEncounters []struct {
        Pokemon struct {
            Name string `json:"name"`
            URL  string `json:"url"`
        } `json:"pokemon"`
        VersionDetails []struct {
            Version struct {
                Name string `json:"name"`
            } `json:"version"`
            MaxChance int `json:"max_chance"`
        } `json:"version_details"`
    } `json:"pokemon_encounters"`
}

type Pokemon struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	BaseExperience int `json:"base_exp"`
	Height int    `json:"height"`
	Weight int    `json:"weight"`
	Types  []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
}


