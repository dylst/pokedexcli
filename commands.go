package main

import (
	"fmt"
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