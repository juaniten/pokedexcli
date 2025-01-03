package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func repl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		text := scanner.Text()
		firstWord := strings.Fields(strings.ToLower(text))[0]

		commands := getCommands()
		if command, ok := commands[firstWord]; ok {
			err := command.callback()
			if err != nil {
				fmt.Println(err.Error())
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Display 20 new locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display 20 previous locations",
			callback:    commandMapb,
		},
	}
}
