package main

import (
	"os"
	"time"

	"github.com/dagregi/pokedexcli/internal/pokeapi"
)

func main() {
	client := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: client,
	}

	replStart(os.Stdin, os.Stdout, cfg)
}
