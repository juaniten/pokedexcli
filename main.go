package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	for {
		fmt.Print("Pokedex > ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		text := scanner.Text()
		cleanInput := strings.Fields(strings.ToLower(text))
		firstWord := cleanInput[0]

		fmt.Println("Your command was:", firstWord)
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
