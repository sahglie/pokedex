package api

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestGetLocationsSuccess(t *testing.T) {
	cli := NewClient(true)
	locations, _ := cli.GetLocations(1)
	fmt.Printf("%#v\n", locations)
}

func TestPokemonJson(t *testing.T) {
	printEndpointJSON("pokemon/arbok")
}

func TestGetArea(t *testing.T) {
	cli := NewClient(true)

	area, _ := cli.GetArea("canalave-city")

	names := make([]string, 0)
	for _, p := range area.PokemonEncounters {
		names = append(names, p.Pokemon.Name)
	}
	fmt.Printf("%#v\n", names)
}

func TestMarshal_pokemonJSON(t *testing.T) {
	body, _ := loadJsonFile("./pokemon.json")

	pokemon := pokemonJSON{}
	err := json.Unmarshal(body, &pokemon)
	if err != nil {
		fmt.Println(err)
	}

	if pokemon.Name != "arbok" {
		t.Errorf("want %s got %s\n", "arbok", pokemon.Name)
	}

	if pokemon.Height != 35 {
		t.Errorf("want %d got %d\n", 35, pokemon.Height)
	}

	if pokemon.Weight != 650 {
		t.Errorf("want %d got %d\n", 650, pokemon.Weight)
	}

	types := make([]string, 0)
	for _, t := range pokemon.Types {
		types = append(types, t.Type.Name)
	}

	if !reflect.DeepEqual(types, []string{"poison"}) {
		fmt.Printf("want %v got %v\n", types, []string{"poison"})
	}
}

func TestMarshal_areaJSON(t *testing.T) {
	body, _ := loadJsonFile("./area.json")

	area := areaJSON{}
	err := json.Unmarshal(body, &area)
	if err != nil {
		fmt.Println(err)
	}

	if area.Name != "canalave-city-area" {
		t.Errorf("want %s got %s\n", "canalave-city-area", area.Name)
	}

	names := make([]string, 0)
	for _, e := range area.PokemonEncounters {
		names = append(names, e.Pokemon.Name)
	}

	if !reflect.DeepEqual(names, []string{"tentacool", "tentacruel"}) {
		fmt.Printf("want %v got %v\n", names, []string{"tentacool", "tentacruel"})
	}
}

func loadJsonFile(name string) ([]byte, error) {
	fd, err := os.Open(name)

	if err != nil {
		return []byte{}, err
	}

	defer fd.Close()

	s := bufio.NewScanner(fd)

	lines := make([]string, 0)
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	body := []byte(strings.Join(lines, "\n"))
	return body, nil
}

func printEndpointJSON(endpoint string) {
	locUrl := getUrl(endpoint, url.Values{})
	cli := NewClient(true)
	data, _ := cli.httpRequest(locUrl)
	fmt.Println(string(data))
}
