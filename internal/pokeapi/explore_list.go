package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ExploreLocations(location string) (RespExplore, error) {
	url := baseURL + "/location-area/" + location

	if val, ok := c.cache.Get(url); ok {
		locationsResp := RespExplore{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return RespExplore{}, err
		}
		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespExplore{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespExplore{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespExplore{}, err
	}

	locationsResp := RespExplore{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespExplore{}, err
	}

	c.cache.Add(url, dat)
	return locationsResp, nil
}
