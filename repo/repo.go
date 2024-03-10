package repo

import (
	"github.com/sahglie/pokedex/api"
)

type Repo struct {
	client api.Client
}

func NewRepo() Repo {
	return Repo{
		client: api.NewClient(),
	}
}

func (r *Repo) Locations(page int) ([]string, error) {
	locations, _ := r.client.GetLocations(page)

	names := make([]string, len(locations))
	for i, l := range locations {
		names[i] = l.Name
	}

	return names, nil
}
