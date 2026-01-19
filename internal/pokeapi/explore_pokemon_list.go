package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// List Pokemon in the area -

func (c *Client) ListPokemon(area string) (RespShallowPokemon, error) {
	url := baseURL + "/location-area/" + area

	if val, ok := c.cache.Get(url); ok {
		pokemonResp := RespShallowPokemon{}
		err := json.Unmarshal(val, &pokemonResp)
		if err != nil {
			return RespShallowPokemon{}, err
		}
		fmt.Println("CACHING SUCCES")
		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowPokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowPokemon{}, nil
	}

	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowPokemon{}, nil
	}
	pokemonResp := RespShallowPokemon{}
	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return RespShallowPokemon{}, err
	}

	c.cache.Add(url, dat)
	return pokemonResp, nil
}
