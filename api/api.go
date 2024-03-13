package api

import (
	"encoding/json"
	"fmt"
	pokecache "github.com/sahglie/pokedex/cache"
	"io"
	"net/http"
	"net/url"
	"time"
)

var apiRoot = "https://pokeapi.co"

type Client interface {
	GetLocations(int) (locationsJSON, error)
	GetArea(string) (areaJSON, error)
	GetPokemon(string) (pokemonJSON, error)
	GET(string, url.Values) ([]byte, error)
}

type PokemonClient struct {
	baseUrl   string
	debugMode bool
	cache     pokecache.Cache
}

func NewClient(debugMode bool) Client {
	return &PokemonClient{
		baseUrl:   apiRoot,
		debugMode: debugMode,
		cache:     pokecache.NewCache(5 * time.Minute),
	}
}

func (c *PokemonClient) GetLocations(page int) (locationsJSON, error) {
	params := buildPageParams(page)

	body, err := c.GET("location-area", params)
	if err != nil {
		return locationsJSON{}, err
	}

	locations := locationsJSON{}
	err = json.Unmarshal(body, &locations)
	if err != nil {
		return locations, err
	}

	return locations, nil
}

func (c *PokemonClient) GetPokemon(name string) (pokemonJSON, error) {
	endpoint := fmt.Sprintf("%s/%s", "pokemon", name)

	body, err := c.GET(endpoint, url.Values{})
	if err != nil {
		return pokemonJSON{}, err
	}

	pokemon := pokemonJSON{}
	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		return pokemon, err
	}

	return pokemon, nil
}

func (c *PokemonClient) GetArea(name string) (areaJSON, error) {
	endpoint := fmt.Sprintf("%s/%s", "location-area", name)

	body, err := c.GET(endpoint, url.Values{})
	if err != nil {
		return areaJSON{}, err
	}

	area := areaJSON{}
	err = json.Unmarshal(body, &area)
	if err != nil {
		return area, err
	}

	return area, nil
}

func (c *PokemonClient) GET(endpoint string, params url.Values) ([]byte, error) {
	eUrl := buildURL(endpoint, params)
	body, ok := c.cache.Get(eUrl)

	if !ok {
		if c.debugMode {
			fmt.Printf("GET: %s\n", eUrl)
		}

		resp, err := http.Get(eUrl)
		if err != nil {
			return body, err
		}

		defer resp.Body.Close()

		body, err = io.ReadAll(resp.Body)
		if err != nil {
			return body, err
		}

		if resp.StatusCode != http.StatusOK {
			return body, fmt.Errorf("status code: %d", resp.StatusCode)
		}

		c.cache.Add(eUrl, body)
	}

	return body, nil
}

func buildURL(endpoint string, params url.Values) string {
	u, _ := url.ParseRequestURI(apiRoot)
	u.Path = fmt.Sprintf("/api/v2/%s", endpoint)
	u.RawQuery = params.Encode()
	return fmt.Sprintf("%v", u)
}

func buildPageParams(page int) url.Values {
	params := url.Values{}

	if page < 1 {
		page = 1
	}

	if page > 1 {
		params.Add("limit", "20")
		params.Add("offset", fmt.Sprintf("%d", (page-1)*20))
	}

	return params
}
