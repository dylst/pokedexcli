package main

import (
	"fmt"
	"math/rand"
	"os"
)

func commandExit(c *config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config, args ...string) error {
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Println("\nUsage:")
	fmt.Println()

	for _, value := range getCommands() {
		fmt.Printf("%s: %s\n", value.name, value.description)
	}
	fmt.Println()
	return nil
}

func commandMap(c *config, args ...string) error {
	locationAreas, err := c.pokeApiClient.ListLocations(c.Next)
	if err != nil {
		return err
	}

	// fmt.Println(locationAreas)
	for _, loc := range locationAreas.Results {
		fmt.Println(loc.Name)
	}
	c.Next = locationAreas.Next
	c.Previous = locationAreas.Previous
	return nil
}

func commandMapB(c *config, args ...string) error {
	locationAreas, err := c.pokeApiClient.ListLocations(c.Previous)
	if err != nil {
		return err
	}

	// fmt.Println(locationAreas)
	for _, loc := range locationAreas.Results {
		fmt.Println(loc.Name)
	}
	c.Next = locationAreas.Next
	c.Previous = locationAreas.Previous
	return nil
}

func commandExplore(c *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("you must provide a location name")
	}
	name := args[0]
	location, err := c.pokeApiClient.ExploreLocation(name)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", location.Name)
	fmt.Println("Found Pokemon:")
	for _, en := range location.PokemonEncounters {
		fmt.Printf("- %s\n", en.Pokemon.Name)
	}
	return nil
}

func commandCatch(c *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("you must provide a pokemon name")
	}
	name := args[0]
	pokemon, err := c.pokeApiClient.GetPokemon(name)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	success := catchHelper(pokemon.BaseExperience)
	if success {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		fmt.Println("You may now inspect it with the inspect command.")
		c.caughtPokemon[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	return nil
}

func catchHelper(baseExperience int) bool {
	catchRate := (float64(baseExperience) / 1000) * 1.5
	random := rand.Float64()
	return catchRate > random
}

func commandInspect(c *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("you must provide a pokemon name")
	}
	name := args[0]
	pokemon, exists := c.caughtPokemon[name]; 
	if !exists {
		return fmt.Errorf("you have not caught that pokemon")
	}
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("-%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("- %s\n", t.Type.Name)
	}
	return nil
}

func commandPokedex(c *config, args ...string) error {
	fmt.Println("Your pokedex:")

	for _, pkmn := range c.caughtPokemon {
		fmt.Printf("- %s\n", pkmn.Name)
	}
	return nil
}