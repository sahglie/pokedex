package main

import (
	"bufio"
	"fmt"
	"github.com/sahglie/pokedex/cmd"
	"github.com/sahglie/pokedex/config"
	"os"
	"strings"
)

var cmds = map[string]cmd.Command{
	"help": {
		Name:        "help",
		Description: "Displays a help message",
		Fn:          cmd.HelpCmd,
	},
	"exit": {
		Name:        "exit",
		Description: "Exit the Pokedex",
		Fn:          cmd.ExitCmd,
	},
	"map": {
		Name:        "map",
		Description: "Displays the names of 20 location areas",
		Fn:          cmd.MapCmd,
	},
	"mapb": {
		Name:        "mapb",
		Description: "Displays the names of previous 20 location areas",
		Fn:          cmd.MapbCmd,
	},
}

func main() {
	cnf := config.NewAppConfig()
	fmt.Printf("%s", cnf.Prompt)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		name := strings.Trim(line, " ")
		command, ok := cmds[name]

		if name == "exit" {
			command.Fn(cnf)
			os.Exit(0)
		}

		if !ok {
			fmt.Println("unknown command")
			cmds["help"].Fn(cnf)
			fmt.Printf("%s", cnf.Prompt)
			continue
		}

		command.Fn(cnf)

		fmt.Printf("%s", cnf.Prompt)
	}
}
