package pokeapi

func (c *Client) Inspect(pokemon *string) Pokemon {
	elem, ok := caughtPokemon[*pokemon]
	if !ok {
		return Pokemon{}
	}
	return elem
}