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
	callback    func(*config, ...string) error
}

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	pokedex          map[string]pokeapi.Pokemon
}

func replStart(in io.Reader, out io.Writer, cfg *config) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Fprint(out, "\x1B[1;2mPokedex >\x1B[0m ")
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		fields := strings.Fields(strings.ToLower(scanner.Text()))
		commandName := fields[0]
		parameters := fields[1:]

		command, ok := getCommands()[commandName]
		if ok {
			err := command.callback(cfg, parameters...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Printf("\x1B[1;31mError:\x1B[0m no such command: %s\n", commandName)
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
		"catch": {
			name:        "catch <pokemon>",
			description: "Try to catch the given pokemon",
			callback:    catchCommand,
		},
		"explore": {
			name:        "explore <area>",
			description: "Get the list of Pokemon in an area",
			callback:    exploreCommand,
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
