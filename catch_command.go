package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func catchCommand(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("missing pokemon name")
	}

	pokemon, err := cfg.pokeapiClient.GetPokemon(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	if rand.Intn(pokemon.BaseExperience) > 40 {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemon.Name)
	fmt.Printf("You may now inspect it with the \x1B[1;2minspect\x1B[0m command\n")
	cfg.pokedex[pokemon.Name] = pokemon

	return nil
}
