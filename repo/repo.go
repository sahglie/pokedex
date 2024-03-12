package repo

import (
	"fmt"
	"github.com/sahglie/pokedex/api"
)

type Repo struct {
	client api.Client
	cache  map[string][]string
	page   int
}

func NewRepo() Repo {
	c := api.NewClient(true)
	return Repo{
		client: c,
		cache:  make(map[string][]string),
		page:   1,
	}
}

func (r *Repo) LocationsNext() ([]string, error) {
	names, _ := r.locations(r.page)
	r.page++
	return names, nil
}

func (r *Repo) LocationsPrev() ([]string, error) {
	if r.page > 1 {
		r.page--
	}

	names, _ := r.locations(r.page)
	return names, nil
}

func (r *Repo) locations(page int) ([]string, error) {
	key := fmt.Sprintf("locations-%d", page)

	names, ok := r.cache[key]

	if !ok {
		locations, _ := r.client.GetLocations(page)

		names = make([]string, len(locations))
		for i, l := range locations {
			names[i] = l.Name
		}

		r.cache[key] = names
	}

	return names, nil
}
