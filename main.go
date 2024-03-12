package main

import (
	"bufio"
	"fmt"
	"github.com/sahglie/pokedex/cmd"
	"github.com/sahglie/pokedex/config"
	"os"
	"strings"
)

func main() {
	cnf := config.NewAppConfig()
	cmds := cmd.Commands

	fmt.Printf("%s", cnf.Prompt)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		name := strings.Trim(line, " ")
		command, ok := cmds[name]

		if name == "exit" || name == "quit" {
			command.Fn(&cnf)
			os.Exit(0)
		}

		if !ok {
			fmt.Println("unknown command")
			cmds["help"].Fn(&cnf)
			fmt.Printf("%s", cnf.Prompt)
			continue
		}

		command.Fn(&cnf)

		fmt.Printf("%s", cnf.Prompt)
	}
}
