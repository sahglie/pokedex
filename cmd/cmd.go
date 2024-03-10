package cmd

import (
	"fmt"
	"github.com/sahglie/pokedex/config"
	"github.com/sahglie/pokedex/repo"
)

type Command struct {
	Name        string
	Description string
	Fn          func(config.AppConfig) error
}

func HelpCmd(c config.AppConfig) error {
	help := `
Welcome to the pokedex!
Usage:

help: Displays a help message
exit: Exit the pokedex
`
	fmt.Println(help)
	return nil
}

func ExitCmd(c config.AppConfig) error {
	fmt.Println("exit")
	return nil
}

func MapCmd(c config.AppConfig) error {
	r := repo.NewRepo()

	names, _ := r.Locations(c.NextPage)
	c.AdvancePager()

	for _, n := range names {
		fmt.Println(n)
	}

	return nil
}

func MapbCmd(c config.AppConfig) error {
	r := repo.NewRepo()

	names, _ := r.Locations(c.PrevPage)
	c.RewindPager()

	for _, n := range names {
		fmt.Println(n)
	}

	return nil
}
