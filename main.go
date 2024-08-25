package main

import (
	"os"
	"time"

	"github.com/dagregi/pokedexcli/internal/pokeapi"
)

func main() {
	client := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{
		pokedex:       make(map[string]pokeapi.Pokemon),
		pokeapiClient: client,
	}

	replStart(os.Stdin, os.Stdout, cfg)
}
