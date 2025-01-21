package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	endpoint := "/pokemon/" + pokemonName
	fullURL := baseURL + endpoint

	// check the cache
	data, ok := c.cache.Get(fullURL)
	if ok {
		fmt.Println("cache hit!")
		pokemon := Pokemon{}
		err := json.Unmarshal(data, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil
	}

	fmt.Println("cache miss!")

	request, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return Pokemon{}, err
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return Pokemon{}, err
	}

	defer response.Body.Close()

	if response.StatusCode > 399 {
		return Pokemon{}, fmt.Errorf("bad status code: %v", response.StatusCode)
	}

	data, err = io.ReadAll(response.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemon := Pokemon{}
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(fullURL, data)
	return pokemon, nil
}