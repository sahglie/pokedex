package config

import (
	"fmt"
	"github.com/sahglie/pokedex/repo"
)

type AppConfig struct {
	Repo      *repo.Repo
	Prompt    string
	DebugMode bool
}

var reset = "\033[0m"
var green = "\033[32m"

func NewAppConfig() AppConfig {
	r := repo.NewRepo()

	return AppConfig{
		Repo:      &r,
		Prompt:    fmt.Sprintf("%sPokedex > %s", green, reset),
		DebugMode: false,
	}
}
