package main

import (
	"fmt"
	"os"
)

func commandExit(c *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config) error {
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Println("\nUsage:")
	fmt.Println()

	for _, value := range getCommands() {
		fmt.Printf("%s: %s\n", value.name, value.description)
	}
	fmt.Println()
	return nil
}

func commandMap(c *config) error {
	locationAreas, err := c.pokeApiClient.ListLocations(c.Next, c.pokeCache)
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

func commandMapB(c *config) error {
	locationAreas, err := c.pokeApiClient.ListLocations(c.Previous, c.pokeCache)
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