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
	client   api.Client
	cache    map[string][]string
	nextPage int
	prevPage int
}

func NewRepo() Repo {
	c := api.NewClient(true)
	return Repo{
		client:   c,
		cache:    make(map[string][]string),
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

func (r *Repo) pages() [2]int {
	return [2]int{r.nextPage, r.prevPage}
}
