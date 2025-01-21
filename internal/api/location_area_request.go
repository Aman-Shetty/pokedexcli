package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResponse, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint

	if pageURL != nil {
		fullURL = *pageURL
	}

	// check the cache
	data, ok := c.cache.Get(fullURL)
	if ok {
		fmt.Println("cache hit!")
		locationAreasResponse := LocationAreasResponse{}
		err := json.Unmarshal(data, &locationAreasResponse)
		if err != nil {
			return LocationAreasResponse{}, err
		}
		return locationAreasResponse, nil
	}

	fmt.Println("cache miss!")

	request, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	defer response.Body.Close()

	if response.StatusCode > 399 {
		return LocationAreasResponse{}, fmt.Errorf("bad status code: %v", response.StatusCode)
	}

	data, err = io.ReadAll(response.Body)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	locationAreasResponse := LocationAreasResponse{}
	err = json.Unmarshal(data, &locationAreasResponse)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	c.cache.Add(fullURL, data)
	return locationAreasResponse, nil
}

func (c *Client) GetLocationArea(LocationAreaName string) (LocationArea, error) {
	endpoint := "/location-area/" + LocationAreaName
	fullURL := baseURL + endpoint

	// check the cache
	data, ok := c.cache.Get(fullURL)
	if ok {
		fmt.Println("cache hit!")
		locationArea := LocationArea{}
		err := json.Unmarshal(data, &locationArea)
		if err != nil {
			return LocationArea{}, err
		}
		return locationArea, nil
	}

	fmt.Println("cache miss!")

	request, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationArea{}, err
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return LocationArea{}, err
	}

	defer response.Body.Close()

	if response.StatusCode > 399 {
		return LocationArea{}, fmt.Errorf("bad status code: %v", response.StatusCode)
	}

	data, err = io.ReadAll(response.Body)
	if err != nil {
		return LocationArea{}, err
	}

	locationArea := LocationArea{}
	err = json.Unmarshal(data, &locationArea)
	if err != nil {
		return LocationArea{}, err
	}

	c.cache.Add(fullURL, data)
	return locationArea, nil
}