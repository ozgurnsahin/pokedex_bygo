package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/ozgurnsahin/pokedex_bygo/internal/pokecache"
)

// ListLocations -
func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	cache := pokecache.NewCache(5 * time.Millisecond)
	url := baseURL + "/location-area"

	if pageURL != nil {
		url = *pageURL
	}
	
	if elem, ok := cache.Entry[url]; !ok {

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return RespShallowLocations{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return RespShallowLocations{}, err
		}
		defer resp.Body.Close()

		dat, err := io.ReadAll(resp.Body)
		if err != nil {
			return RespShallowLocations{}, err
		}

		cache.Add(url,dat)

		locationsResp := RespShallowLocations{}
		err = json.Unmarshal(dat, &locationsResp)
		if err != nil {
			return RespShallowLocations{}, err
		}

		return locationsResp, nil

	} else {

		locationsResp := RespShallowLocations{}
		err := json.Unmarshal(elem.Val, &locationsResp)

		if err != nil {
			return RespShallowLocations{}, err
		}

		return locationsResp, nil
	}
	
}

