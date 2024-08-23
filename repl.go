package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func replStart(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Fprint(out, "Pokedex > ")
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		fields := strings.Fields(scanner.Text())
		commandName := fields[0]

		command, ok := getCommands()[commandName]
		if ok {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Printf("No such command: %s\n", commandName)
			continue
		}
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Print this help text and exit",
			callback:    helpCommand,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    exitCommand,
		},
	}
}
