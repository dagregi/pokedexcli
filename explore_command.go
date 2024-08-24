package main

import (
	"errors"
	"fmt"
)

func exploreCommand(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("missing area name")
	}

	location, err := cfg.pokeapiClient.GetLocation(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", args[0])
	fmt.Println("Found Pokemon:")
	for _, pokemonEnc := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemonEnc.Pokemon.Name)
	}

	return nil
}
