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
		line = strings.Trim(line, " ")
		tokens := strings.Split(line, " ")

		name := tokens[0]

		args := make([]string, 0)
		if len(tokens) > 1 {
			args = tokens[1:]
		}

		command, ok := cmds[name]

		if !ok {
			fmt.Println("unknown command")
			cmds["help"].Fn(&cnf)
			fmt.Printf("%s", cnf.Prompt)
			continue
		}

		err := command.Fn(&cnf, args...)
		if err != nil {
			fmt.Println(err)
		}

		if name == "exit" || name == "quit" {
			os.Exit(0)
		}

		fmt.Printf("%s", cnf.Prompt)
	}
}
