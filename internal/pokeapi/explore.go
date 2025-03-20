package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/ozgurnsahin/pokedex_bygo/internal/pokecache"
)

// Explore Locations -
func (c *Client) Explore(area *string) (RespExploreMap, error) {
	cache := pokecache.NewCache(5 * time.Second)
	url := baseURL + "/location-area"

	if area != nil {
		url = baseURL + "/location-area/" + *area
	}
	
	if elem, ok := cache.Entry[url]; !ok {

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return RespExploreMap{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return RespExploreMap{}, err
		}
		defer resp.Body.Close()

		dat, err := io.ReadAll(resp.Body)
		if err != nil {
			return RespExploreMap{}, err
		}

		cache.Add(url,dat)

		var areaDetail RespExploreMap
        err = json.Unmarshal(dat, &areaDetail)
        if err != nil {
            return RespExploreMap{}, err
        }
		
        return areaDetail, nil

	} else {

		var areaDetail RespExploreMap
        err := json.Unmarshal(elem.Val, &areaDetail)
        if err != nil {
            return RespExploreMap{}, err
        }
		
        return areaDetail, nil
	}
	
}

