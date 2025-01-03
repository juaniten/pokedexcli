package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getLocations(url string) (LocationAreas, error) {
	res, err := http.Get(url)
	if err != nil {
		return LocationAreas{}, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return LocationAreas{}, fmt.Errorf("response failed with status code: %d", res.StatusCode)
	}

	var la LocationAreas
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&la); err != nil {
		fmt.Println("error decoding response body")
		return LocationAreas{}, err
	}

	return la, nil

}

func printLocations(la LocationAreas) {

	locations := la.Result
	names := make([]string, len(locations))
	for i, location := range locations {
		names[i] = location.Name
	}

	for _, name := range names {
		fmt.Println(name)
	}
}

func commandMap(config *Config) error {
	if config.Next == "" {
		fmt.Println("there are no more results")
		return nil
	}
	la, err := getLocations(config.Next)
	if err != nil {
		return err
	}

	printLocations(la)

	config.Previous = la.Previous
	config.Next = la.Next

	return nil
}

func commandMapb(config *Config) error {
	if config.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	la, err := getLocations(config.Previous)
	if err != nil {
		return err
	}

	printLocations(la)

	config.Previous = la.Previous
	config.Next = la.Next
	return nil
}

type Location struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type LocationAreas struct {
	Count    int        `json:"count"`
	Next     string     `json:"next"`
	Previous string     `json:"previous"`
	Result   []Location `json:"results"`
}
