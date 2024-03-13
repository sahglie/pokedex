package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

var apiRoot = "https://pokeapi.co"

type Client interface {
	GetLocations(int) ([]Location, error)
	GetArea(string) (areaJSON, error)
	GetPokemon(string) (pokemonJSON, error)
	httpRequest(string) ([]byte, error)
}

type Location struct {
	Name string
	Url  string
}

type Area struct {
	ID      int
	Name    string
	Pokemon []string
}

type PokemonClient struct {
	baseUrl   string
	debugMode bool
}

func NewClient(debugMode bool) Client {
	return &PokemonClient{
		baseUrl:   apiRoot,
		debugMode: debugMode,
	}
}

func (c *PokemonClient) GetLocations(page int) ([]Location, error) {
	locations := make([]Location, 0)

	params := buildParams(page)
	locUrl := getUrl("location-area", params)

	body, err := c.httpRequest(locUrl)

	if err != nil {
		return locations, err
	}

	jsonResponse := locationsJSON{}
	err = json.Unmarshal(body, &jsonResponse)
	if err != nil {
		return locations, err
	}

	locations = jsonResponse.Results
	return locations, nil
}

func (c *PokemonClient) GetPokemon(name string) (pokemonJSON, error) {
	endpoint := fmt.Sprintf("%s/%s", "pokemon", name)
	locUrl := getUrl(endpoint, url.Values{})

	body, err := c.httpRequest(locUrl)
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
	locUrl := getUrl(endpoint, url.Values{})

	body, err := c.httpRequest(locUrl)
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

func (c *PokemonClient) httpRequest(url string) ([]byte, error) {
	var body []byte

	if c.debugMode {
		fmt.Printf("GET: %s\n", url)
	}

	resp, err := http.Get(url)
	if err != nil {
		return body, err
	}

	if resp.StatusCode != http.StatusOK {
		return body, fmt.Errorf("status code: %d", resp.StatusCode)
	}

	defer resp.Body.Close()
	body, err = io.ReadAll(resp.Body)

	if err != nil {
		return body, err
	}

	return body, nil
}

func getUrl(endpoint string, params url.Values) string {
	u, _ := url.ParseRequestURI(apiRoot)
	u.Path = fmt.Sprintf("/api/v2/%s", endpoint)
	u.RawQuery = params.Encode()
	return fmt.Sprintf("%v", u)
}

func buildParams(page int) url.Values {
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
