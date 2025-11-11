package main

import (
	"time"

	"github.com/dylst/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(time.Second * 5, time.Minute * 5)
	cfg := config{
		pokeApiClient: pokeClient,
	}
	startRepl(cfg)
}