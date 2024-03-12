package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

var apiRoot = "https://pokeapi.co"

type jsonData struct {
	Count   int
	NextUrl string `json:"next"`
	PrevUrl string `json:"previous"`
	Results []Location
}

type Client interface {
	GetLocations(int) ([]Location, error)
}

type Location struct {
	Name string
	Url  string
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
	locUrl := getUrl("location", params)

	body, err := c.httpRequest(locUrl)

	if err != nil {
		return locations, err
	}

	jsonResponse := jsonData{}
	err = json.Unmarshal(body, &jsonResponse)
	if err != nil {
		return locations, err
	}

	locations = jsonResponse.Results
	return locations, nil
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
