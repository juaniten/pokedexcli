package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func main() {
	commands := map[string]cliCommand{}

	commands["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the Pokedex",
		callback: func() error {
			return commandExit()
		},
	}
	commands["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback: func() error {
			return commandHelp(commands)
		},
	}
	commands["map"] = cliCommand{
		name:        "map",
		description: "Display 20 new locations",
		callback: func() error {
			return commandMap()
		},
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		scanner.Scan()
		text := scanner.Text()
		firstWord := strings.Fields(strings.ToLower(text))[0]

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

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(commands map[string]cliCommand) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for name, command := range commands {
		fmt.Printf("%s: %s\n", name, command.description)
	}
	return nil
}

func commandMap() error {
	return nil
}
