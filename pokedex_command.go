package main

import "fmt"

func pokedexCommand(cfg *config, args ...string) error {
	if len(cfg.pokedex) == 0 {
		fmt.Println("You haven't caught any pokemons!")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.pokedex {
		fmt.Printf(" - %s\n", pokemon.Name)
	}
	return nil
}
