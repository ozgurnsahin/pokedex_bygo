package main


func commandCatch(cfg *config) error {

	_, err := cfg.pokeapiClient.Catch(cfg.pokemonName)

	if err != nil {
		return err
	}

	return nil
}
