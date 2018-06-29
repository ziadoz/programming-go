package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// A struct with field tags that determine how to parse as JSON.
type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Colour bool `json:"colour,omitempty"`
	Actors []string
}

func main() {
	var movies = []Movie{
		{
			Title:  "Casablanca",
			Year:   1942,
			Colour: false,
			Actors: []string{
				"Humphrey Bogart",
				"Ingrid Bergman",
			},
		},
		{
			Title:  "Cool Hand Luke",
			Year:   1967,
			Colour: true,
			Actors: []string{
				"Paul Newman",
			},
		},
		{
			Title:  "Bullitt",
			Year:   1968,
			Colour: true,
			Actors: []string{
				"Steve McQueen",
				"Jacqueline Bisset",
			},
		},
	}

	// Marshal the movies struct into a JSON byte slice without any identation.
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON marshaling failed %s", err)
	}
	fmt.Printf("%s\n", data)

	// Marshal the movies into an idented, human readable JSON byte slice.
	datai, err := json.MarshalIndent(movies, "", "    ") // Other params are line prefixed and indentation.
	if err != nil {
		log.Fatalf("JSON marshaling failed %s", err)
	}
	fmt.Printf("%s\n", datai)

	// Unmarshal only the titles from the JSON we just marshaled.
	var titles []struct{ Title string }
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON unmarshaling has failed: %s", err)
	}
	fmt.Println(titles)
}
