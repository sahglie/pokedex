package api

type locationsJSON struct {
	Count   int
	NextUrl string `json:"next"`
	PrevUrl string `json:"previous"`
	Results []Location
}

type areaJSON struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

type pokemonJSON struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Height    int    `json:"height"`
	Weight    int    `json:"weight"`
	Abilities []struct {
		Ability struct {
			Name string `json:"name"`
		} `json:"ability"`
	} `json:"abilities"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}
