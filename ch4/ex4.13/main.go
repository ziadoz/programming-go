// Usage: export OMDB_API_KEY="" && go run main.go [search query]
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopl.io/ch4/ex4.13/omdb"
)

func main() {
	apikey := os.Getenv("OMDB_API_KEY")

	if apikey == "" {
		fmt.Println("OMDB_API_KEY is not set")
		os.Exit(1)
	}

	query := strings.Join(os.Args[1:], " ")
	if strings.Trim(query, " ") == "" {
		fmt.Println("No search query entered")
		os.Exit(1)
	}

	movie, err := omdb.SearchPosterAPI(omdb.APIKey(apikey), query)
	if err != nil {
		fmt.Printf("Could not get movie information: %s", err)
		os.Exit(1)
	}

	filename := movie.Title + filepath.Ext(movie.Poster)
	if err := movie.DownloadPoster(filename); err != nil {
		fmt.Printf("Could not download movie poster: %s", err)
		os.Exit(1)
	}

	fmt.Println(filename)
}
