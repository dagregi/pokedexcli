package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/dagregi/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

func replStart(in io.Reader, out io.Writer, cfg *config) {
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
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Printf("Error: no such command: %s\n", commandName)
			continue
		}
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Print this help text",
			callback:    helpCommand,
		},
		"map": {
			name:        "map",
			description: "Print the next page of locations",
			callback:    mapfCommand,
		},
		"mapb": {
			name:        "mapb",
			description: "Print the previous page of locations",
			callback:    mapbCommand,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    exitCommand,
		},
	}
}
