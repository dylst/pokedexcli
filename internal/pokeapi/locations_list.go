package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c Client) ListLocations(pageUrl *string) (RespLocations, error) {
	url := baseURL + "/location-area"

	if pageUrl != nil {
		url = *pageUrl
	}

	val, exists := c.cache.Get(url)
	if exists {
		locationsResp := RespLocations{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return RespLocations{}, err
		}
		return locationsResp, nil
	}
	
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
	
}