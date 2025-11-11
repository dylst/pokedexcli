package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c Client) ExploreLocation(endpoint string) (ExploreLocationResponse, error) {
	url := baseURL + "/location-area/" + endpoint

	if val, exists := c.cache.Get(url); exists {
		exploreLocationsResp := ExploreLocationResponse{}
		err := json.Unmarshal(val, &exploreLocationsResp)
		if err != nil {
			return ExploreLocationResponse{}, err
		}
		return exploreLocationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ExploreLocationResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return ExploreLocationResponse{}, err
	}
	defer resp.Body.Close()

	var exploreLocationResp ExploreLocationResponse 
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return ExploreLocationResponse{}, err
	}
	err = json.Unmarshal(data, &exploreLocationResp)
	if err != nil {
		return ExploreLocationResponse{}, err
	}

	c.cache.Add(url, data)
	return exploreLocationResp, nil
}