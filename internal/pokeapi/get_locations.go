package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetLocations(url string) (LocationAreas, error) {
	res, err := http.Get(url)
	if err != nil {
		return LocationAreas{}, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode > 298 {
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

func PrintLocations(la LocationAreas) {

	locations := la.Result
	names := make([]string, len(locations))
	for i, location := range locations {
		names[i] = location.Name
	}

	for _, name := range names {
		fmt.Println(name)
	}
}
