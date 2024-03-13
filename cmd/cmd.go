package cmd

import (
	"errors"
	"fmt"
	"github.com/sahglie/pokedex/config"
	"github.com/sahglie/pokedex/repo"
)

type Command struct {
	Name        string
	Description string
	Fn          func(*config.AppConfig, ...string) error
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
	"explore": {
		Name:        "explore",
		Description: "Displays the names of pokemon in the area",
		Fn:          ExploreArea,
	},
	"catch": {
		Name:        "catch",
		Description: "Try to catch a pokemon in the area",
		Fn:          CatchPokemon,
	},
	"inspect": {
		Name:        "catch",
		Description: "Inspect stats for a pokemon you have caught",
		Fn:          InspectPokemon,
	},
}

func HelpCmd(c *config.AppConfig, args ...string) error {
	help := `Welcome to the pokedex!
Usage:
  help: Displays a help message
  exit: Exit the Pokedex
  quit: Exit the Pokedex
  map: Displays the names of 20 location areas
  mapb: Displays the names of the previous 20 location areas
  explore: Displays the names of pokemon in the area 
`
	fmt.Print(help)

	return nil
}

func ExitCmd(c *config.AppConfig, args ...string) error {
	fmt.Println("bye!")
	return nil
}

func MapCmd(c *config.AppConfig, args ...string) error {
	names, _ := c.Repo.LocationsNext()

	for _, n := range names {
		fmt.Println(n)
	}

	return nil
}

func MapbCmd(c *config.AppConfig, args ...string) error {
	names, err := c.Repo.LocationsPrev()

	if errors.Is(err, repo.ErrNoPrevPage) {
		fmt.Printf("%v\n", err)
		return nil
	}

	if err != nil {
		return err
	}

	for _, n := range names {
		fmt.Println(n)
	}

	return nil
}

func ExploreArea(c *config.AppConfig, args ...string) error {
	if len(args) == 0 {
		return errors.New("`explore <name>` expects a location name")
	}

	fmt.Printf("Exploring %s...\n", args[0])
	fmt.Println("Found pokemon:")
	names, _ := c.Repo.ListPokemonInArea(args[0])
	for _, n := range names {
		fmt.Printf(" - %s\n", n)
	}

	return nil
}

func CatchPokemon(c *config.AppConfig, args ...string) error {
	if len(args) == 0 {
		return errors.New("`catch <name>` expects a pokemon name")
	}

	name := args[0]

	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	fmt.Printf("%s was caught!\n", name)

	pokemon, err := c.Repo.GetPokemonInfo(name)
	if err != nil {
		return err
	}

	fmt.Println(pokemon.String())

	return nil
}

func InspectPokemon(c *config.AppConfig, args ...string) error {
	if len(args) == 0 {
		return errors.New("`inspect <name>` expects a pokemon name")
	}

	name := args[0]

	fmt.Printf("Inspecting %s...\n", name)

	pokemon, err := c.Repo.GetPokemonInfo(name)
	if err != nil {
		return err
	}

	fmt.Println(pokemon.String())

	return nil
}
