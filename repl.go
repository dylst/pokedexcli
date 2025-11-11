package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/dylst/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name string
	description string 
	callback func(*config) error
}

type config struct {
	pokeApiClient pokeapi.Client
	Next *string 
	Previous *string
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"map": {
			name: "map",
			description: "List the next 20 areas",
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "List the previous 20 areas",
			callback: commandMapB,
		},
	}
}

func startRepl(cfg config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex> ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}
		commandName := words[0]

		command, ok := getCommands()[commandName]
		
		if ok {
			err := command.callback(&cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	if text == "" {
		return []string{""}
	}
	lowered := strings.ToLower(text)
	fields := strings.Fields(lowered)
	return fields
}