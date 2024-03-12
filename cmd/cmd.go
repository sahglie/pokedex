package cmd

import (
	"fmt"
	"github.com/sahglie/pokedex/config"
)

type Command struct {
	Name        string
	Description string
	Fn          func(*config.AppConfig) error
}

var Commands = map[string]Command{
	"help": {
		Name:        "help",
		Description: "Displays a help message",
		Fn:          HelpCmd,
	},
	"exit": {
		Name:        "exit",
		Description: "Exit the Pokedex",
		Fn:          ExitCmd,
	},
	"quit": {
		Name:        "exit",
		Description: "Exit the Pokedex",
		Fn:          ExitCmd,
	},
	"map": {
		Name:        "map",
		Description: "Displays the names of 20 location areas",
		Fn:          MapCmd,
	},
	"mapb": {
		Name:        "mapb",
		Description: "Displays the names of the previous 20 location areas",
		Fn:          MapbCmd,
	},
}

func HelpCmd(c *config.AppConfig) error {
	help := `Welcome to the pokedex!
Usage:
  help: Displays a help message
  exit: Exit the Pokedex
  quit: Exit the Pokedex
  map: Displays the names of 20 location areas
  mapb: Displays the names of the previous 20 location areas
`
	fmt.Print(help)

	return nil
}

func ExitCmd(c *config.AppConfig) error {
	fmt.Println("bye!")
	return nil
}

func MapCmd(c *config.AppConfig) error {
	names, _ := c.Repo.LocationsNext()

	for _, n := range names {
		fmt.Println(n)
	}

	return nil
}

func MapbCmd(c *config.AppConfig) error {
	names, _ := c.Repo.LocationsPrev()

	for _, n := range names {
		fmt.Println(n)
	}

	return nil
}
