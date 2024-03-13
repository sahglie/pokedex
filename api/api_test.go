package api

import (
	"bufio"
	"encoding/json"
	"fmt"
	pokecache "github.com/sahglie/pokedex/cache"
	"net/url"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestGET_CachesResponse(t *testing.T) {
	cache := pokecache.NewCache(5 * time.Minute)

	cli := PokemonClient{cache: cache, baseUrl: apiRoot, debugMode: true}

	if cache.Size() != 0 {
		t.Errorf("want %d got %d\n", 0, cache.Size())
	}

	if _, err := cli.GET("location-area", url.Values{}); err != nil {
		fmt.Println(err)
	}

	if cache.Size() != 1 {
		t.Errorf("want %d got %d\n", 1, cache.Size())
	}
}

func TestGetLocations(t *testing.T) {
	cli := NewClient(false)
	locations, _ := cli.GetLocations(1)

	if locations.Count != 1054 {
		t.Errorf("want %d got %d\n", 1054, locations.Count)
	}

	results := locations.Results
	if len(results) != 20 {
		t.Errorf("want %d got %d\n", 20, len(results))
	}

	if results[0].Name != "canalave-city-area" {
		t.Errorf("want %s got %s\n", "canalave-city-area", results[0].Name)
	}

	//fmt.Printf("%v\n", locations)
}

func TestGetArea(t *testing.T) {
	cli := NewClient(false)
	area, _ := cli.GetArea("canalave-city-area")
	pokemon := area.PokemonEncounters

	if len(pokemon) != 11 {
		t.Errorf("want %d got %d\n", 11, len(pokemon))
	}

	p := pokemon[0].Pokemon
	if p.Name != "tentacool" {
		t.Errorf("want %s got %s\n", "tentacool", p.Name)
	}

	//fmt.Printf("%v\n", pokemon)
}

func TestGetPokemon(t *testing.T) {
	cli := NewClient(false)
	pokemon, _ := cli.GetPokemon("tentacool")

	if pokemon.Name != "tentacool" {
		t.Errorf("want %s got %s\n", "tentacool", pokemon.Name)
	}

	if pokemon.Weight != 455 {
		t.Errorf("want %d got %d\n", 455, pokemon.Weight)
	}

	if pokemon.Height != 9 {
		t.Errorf("want %d got %d\n", 9, pokemon.Height)
	}

	stats := pokemon.Stats
	if len(stats) != 6 {
		t.Errorf("want %d got %d\n", 6, len(stats))
	}

	bs := stats[0].BaseStat
	if bs != 40 {
		t.Errorf("want %d got %d\n", 40, bs)
	}

	s := stats[0].Stat.Name
	if s != "hp" {
		t.Errorf("want %s got %s\n", "hp", s)
	}

	//fmt.Printf("%v\n", pokemon)
}

func TestMarshal_areaJSON(t *testing.T) {
	body, _ := loadJsonFile("./area.json")

	area := areaJSON{}
	err := json.Unmarshal(body, &area)
	if err != nil {
		fmt.Println(err)
	}

	pokemons := area.PokemonEncounters
	if len(pokemons) != 2 {
		t.Errorf("want %d got %d\n", 2, len(area.PokemonEncounters))
	}

	p := pokemons[0].Pokemon.Name
	if p != "tentacool" {
		t.Errorf("want %s got %s\n", "tentacool", p)
	}

	//fmt.Printf("%v\n", area)
}

func TestMarshal_locationsJSON(t *testing.T) {
	body, _ := loadJsonFile("./location-areas.json")

	locations := locationsJSON{}
	err := json.Unmarshal(body, &locations)
	if err != nil {
		fmt.Println(err)
	}

	if locations.Count != 1054 {
		t.Errorf("want %d got %d\n", 1054, locations.Count)
	}

	areas := locations.Results
	if len(areas) != 20 {
		t.Errorf("want %d got %d\n", 20, len(areas))
	}

	area := areas[0].Name
	if area != "canalave-city-area" {
		t.Errorf("want %s got %s\n", "canalave-city-area", area)
	}
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

	stats := make([]string, 0)
	for _, t := range pokemon.Stats {
		stats = append(stats, t.Stat.Name)
	}

	if !reflect.DeepEqual(stats, []string{"speed"}) {
		fmt.Printf("want %v got %v\n", types, []string{"speed"})
	}
}

//func TestPokemonJson(t *testing.T) {
//	printEndpointJSON("location-area")
//}

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
	cli := NewClient(true)
	data, err := cli.GET(endpoint, url.Values{})

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))
}
