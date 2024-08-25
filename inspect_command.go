package main

import (
	"errors"
	"fmt"
)

func inspectCommand(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("missing pokemon name")
	}

	pokemonName := args[0]
	if pokemon, ok := cfg.pokedex[pokemonName]; ok {
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)

		fmt.Println("Stats:")
		for _, s := range pokemon.Stats {
			fmt.Printf(" -%s: %d\n", s.Stat.Name, s.BaseStat)
		}

		fmt.Println("Types:")
		for _, t := range pokemon.Types {
			fmt.Printf(" - %s\n", t.Type.Name)
		}

		return nil
	}

	fmt.Println("you have not caught that pokemon")
	return nil
}
