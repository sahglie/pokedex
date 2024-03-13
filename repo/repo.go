package repo

import (
	"errors"
	"fmt"
	"github.com/sahglie/pokedex/api"
)

var (
	ErrNoPrevPage = errors.New("no prev page")
	ErrNoNextPage = errors.New("no next page")
)

type Repo struct {
	client    api.Client
	nextPage  int
	prevPage  int
	pokemonDb PokemonDB
}

type PokemonDB map[string]Pokemon

func NewRepo() Repo {
	c := api.NewClient(false)

	return Repo{
		client:    c,
		nextPage:  1,
		prevPage:  -1,
		pokemonDb: PokemonDB{},
	}
}

func (r *Repo) LocationsNext() ([]string, error) {
	names, err := r.locations(r.nextPage)
	if err != nil {
		return names, err
	}

	r.nextPage++
	r.prevPage++
	return names, nil
}

func (r *Repo) LocationsPrev() ([]string, error) {
	if r.prevPage < 1 {
		return []string{}, ErrNoPrevPage
	}

	names, err := r.locations(r.prevPage)
	if err != nil {
		return names, err
	}

	r.prevPage--
	r.nextPage--

	return names, nil
}

func (r *Repo) ListPokemonInArea(name string) ([]string, error) {
	area, err := r.client.GetArea(name)
	if err != nil {
		return []string{}, err
	}

	names := make([]string, 0)
	for _, e := range area.PokemonEncounters {
		names = append(names, e.Pokemon.Name)
	}

	return names, nil
}

func (r *Repo) GetPokemonInfo(name string) (Pokemon, error) {
	p, err := r.client.GetPokemon(name)
	if err != nil {
		return Pokemon{}, err
	}

	stats := make([]string, 0)
	for _, s := range p.Stats {
		stats = append(stats, fmt.Sprintf("%s: %d", s.Stat.Name, s.BaseStat))
	}

	types := make([]string, 0)
	for _, t := range p.Types {
		types = append(types, t.Type.Name)
	}

	pokemon := Pokemon{
		Id:     p.ID,
		Name:   p.Name,
		Height: p.Height,
		Weight: p.Weight,
		Stats:  stats,
		Types:  types,
	}

	return pokemon, nil
}

func (r *Repo) SavePokemon(p Pokemon) {
	r.pokemonDb[p.Name] = p
}

func (r *Repo) GetSavedPokemon(name string) (Pokemon, bool) {
	p, ok := r.pokemonDb[name]
	if !ok {
		return Pokemon{}, false
	}
	return p, true
}

func (r *Repo) ListSavedPokemons() []Pokemon {
	pokemons := make([]Pokemon, 0)
	for _, v := range r.pokemonDb {
		pokemons = append(pokemons, v)
	}
	return pokemons
}

func (r *Repo) locations(page int) ([]string, error) {
	locations, err := r.client.GetLocations(page)
	if err != nil {
		return []string{}, err
	}

	names := make([]string, 0)
	for _, r := range locations.Results {
		names = append(names, r.Name)
	}

	return names, nil
}

func (r *Repo) pages() [2]int {
	return [2]int{r.nextPage, r.prevPage}
}
