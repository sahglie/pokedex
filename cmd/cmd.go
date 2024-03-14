package cmd

import (
	"errors"
	"fmt"
	"github.com/sahglie/pokedex/config"
	"github.com/sahglie/pokedex/repo"
	"math/rand"
	"time"
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
		Name:        "inspect",
		Description: "Inspect stats for a pokemon you have caught",
		Fn:          InspectPokemon,
	},
	"pokedex": {
		Name:        "pokedex",
		Description: "Inspect stats for a pokemon you have caught",
		Fn:          Pokedex,
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
  catch: Try and catch a pokemon, if you catch them they will be added to your pokedex
  inspect: View stats of a pokemon in your pokedex
  pokedex: View the pokemon in your pokedex
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

	pokemon, err := c.Repo.GetPokemonInfo(name)
	if err != nil {
		return err
	}

	if !attemptCatch(pokemon.BaseExperience) {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	c.Repo.SavePokemon(pokemon)
	fmt.Printf("%s was caught!\n", name)
	fmt.Println("You may now inspect it with the inspect command.")

	return nil
}

func InspectPokemon(c *config.AppConfig, args ...string) error {
	if len(args) == 0 {
		return errors.New("`inspect <name>` expects a pokemon name")
	}

	name := args[0]
	p, ok := c.Repo.GetSavedPokemon(name)

	if !ok {
		fmt.Println("you have not cauth that pokemon")
		return nil
	}

	fmt.Println(p.String())

	return nil
}

func Pokedex(c *config.AppConfig, args ...string) error {
	fmt.Println("Your Pokedex:")

	pokemons := c.Repo.ListSavedPokemons()

	if len(pokemons) == 0 {
		fmt.Println("* pokedex is empty *")
		return nil
	}

	for _, p := range pokemons {
		fmt.Printf("  - %s\n", p.Name)
	}

	return nil
}

func attemptCatch(experience int) bool {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	randVal := r.Float64()
	probabilityThreshold := 1.0 / (1.0 + float64(experience/10))

	return randVal < probabilityThreshold
}
