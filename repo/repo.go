package repo

import (
	"errors"
	"fmt"
	"github.com/sahglie/pokedex/api"
	"github.com/sahglie/pokedex/cache"
	"strings"
	"time"
)

var (
	ErrNoPrevPage = errors.New("no prev page")
	ErrNoNextPage = errors.New("no next page")
)

type Repo struct {
	client   api.Client
	cache    pokecache.Cache
	nextPage int
	prevPage int
}

func NewRepo() Repo {
	c := api.NewClient(true)
	return Repo{
		client:   c,
		cache:    pokecache.NewCache(5 * time.Minute),
		nextPage: 1,
		prevPage: -1,
	}
}

func (r *Repo) LocationsNext() ([]string, error) {
	names, _ := r.locations(r.nextPage)
	r.nextPage++
	r.prevPage++
	return names, nil
}

func (r *Repo) LocationsPrev() ([]string, error) {
	if r.prevPage < 1 {
		return []string{}, ErrNoPrevPage
	}

	names, _ := r.locations(r.prevPage)

	r.prevPage--
	r.nextPage--

	return names, nil
}

func (r *Repo) ListPokemonInArea(name string) ([]string, error) {
	key := fmt.Sprintf("area-%s", name)

	data, ok := r.cache.Get(key)

	if !ok {
		area, _ := r.client.GetArea(name)
		builder := strings.Builder{}
		for _, e := range area.PokemonEncounters {
			builder.WriteString(e.Pokemon.Name)
			builder.WriteString("\n")
		}

		s := builder.String()
		data = []byte(strings.TrimRight(s, "\n"))
		r.cache.Add(key, data)
	}

	names := strings.Split(string(data), "\n")
	return names, nil
}

func (r *Repo) GetPokemonInfo(name string) (Pokemon, error) {
	p, _ := r.client.GetPokemon(name)

	stats := make([]string, 0)
	for _, s := range p.Stats {
		stats = append(stats, fmt.Sprintf("%s - %d", s.Stat.Name, s.BaseStat))
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

func (r *Repo) locations(page int) ([]string, error) {
	key := fmt.Sprintf("locations-%d", page)

	data, ok := r.cache.Get(key)

	if !ok {
		locations, _ := r.client.GetLocations(page)
		builder := strings.Builder{}
		for _, l := range locations {
			builder.WriteString(l.Name)
			builder.WriteString("\n")
		}

		data = []byte(builder.String())
		r.cache.Add(key, data)
	}

	names := strings.Split(string(data), "\n")

	return names, nil
}

func (r *Repo) pages() [2]int {
	return [2]int{r.nextPage, r.prevPage}
}
