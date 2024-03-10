package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

var apiRoot = "https://pokeapi.co"
var DebugMode = false

type Client struct {
	baseUrl   string
	debugMode bool
}

type JsonResponse struct {
	Count   int
	NextUrl string `json:"next"`
	PrevUrl string `json:"previous"`
	Results []Location
}

type Location struct {
	Name string
	Url  string
}

func NewClient() Client {
	return Client{
		baseUrl:   apiRoot,
		debugMode: true,
	}
}

func (c *Client) GetLocations(page int) ([]Location, error) {
	locations := make([]Location, 0)

	params := buildParams(page)
	url := getUrl("location", params)

	body, err := httpRequest(url)

	if err != nil {
		return locations, err
	}

	jsonResponse := JsonResponse{}
	err = json.Unmarshal(body, &jsonResponse)
	if err != nil {
		return locations, err
	}

	if DebugMode {
		fmt.Printf("%+v\n", jsonResponse)
	}

	locations = jsonResponse.Results
	return locations, nil
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
		params.Add("offset", fmt.Sprintf("%d", page*20))
	}

	return params
}

func httpRequest(url string) ([]byte, error) {
	var body []byte

	if DebugMode {
		fmt.Printf("GET: %s\n", url)
	}

	resp, err := http.Get(url)
	if err != nil {
		return body, err
	}

	defer resp.Body.Close()
	body, err = io.ReadAll(resp.Body)

	if err != nil {
		return body, err
	}

	return body, nil
}
