package main

import (
	"time"

	"github.com/dylst/pokedexcli/internal/pokeapi"
	"github.com/dylst/pokedexcli/internal/pokecache"
)

func main() {
	pokeClient := pokeapi.NewClient(time.Second * 60)
	pokeCache := pokecache.NewCache(time.Second * 60)
	cfg := config{
		pokeApiClient: pokeClient,
		pokeCache: pokeCache,
	}
	startRepl(cfg)
}