package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/dylst/pokedexcli/internal/pokecache"
)

func (c Client) ListLocations(pageUrl *string, cache *pokecache.Cache) (RespLocations, error) {
	url := baseURL + "/location-area"

	if pageUrl != nil {
		url = *pageUrl
	}

	cachedResp, exists := cache.Get(url)
	if !exists {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return RespLocations{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return RespLocations{}, nil 
		}
		defer resp.Body.Close()

		data, err := io.ReadAll(resp.Body)
		if err != nil {
			return RespLocations{}, err
		}

		locationsResp := RespLocations{}
		err = json.Unmarshal(data, &locationsResp)
		if err != nil {
			return RespLocations{}, err
		}
		return locationsResp, nil
	} else {
		locationsResp := RespLocations{}
		err := json.Unmarshal(cachedResp, &locationsResp)
		if err != nil {
			return RespLocations{}, err
		}
		return locationsResp, nil
	}
}

/* 
Reformatting my code. This is the response portion. Now we should format the map functions to use this list locations.
*/ 