// Usage: go run main.go -progress -query=php
//        go run main.go -progress -query=python
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"gopl.io/ch4/ex4.12/xkcd"
)

var query string
var progress bool

func init() {
	flag.StringVar(&query, "query", "", "The keywords to search for in the comic cache")
	flag.BoolVar(&progress, "progress", false, "Show comic caching progress")
	flag.Parse()
}

func main() {
	if strings.Trim(query, "") == "" {
		fmt.Fprintln(os.Stderr, "Please enter a search query")
		os.Exit(1)
	}

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not determine current directory %s\n", err)
		os.Exit(1)
	}

	cacheDir := filepath.Join(cwd, "cache")
	if _, err := os.Stat(cacheDir); err != nil {
		fmt.Fprintf(os.Stderr, "Cache directory does not exist: %s\n", err)
		os.Exit(1)
	}

	latest, err := xkcd.GetLatestComicNumber()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not determine latest comic number: %s\n", err)
		os.Exit(1)
	}

	fmt.Println("Download comics to cache…")
	for number := 1; number <= latest; number++ {
		// Skip comic 404, which is a joke.
		if number == 404 {
			continue
		}

		// Skip if the comic JSON file already exists.
		file := filepath.Join(cacheDir, "xkcd-"+strconv.Itoa(number)+".json")
		if _, err := os.Stat(file); err == nil {
			continue
		}

		// Otherwise get the comic and store it in a JSON file.
		comic, err := xkcd.GetComic(number)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to get comic: %s\n", err)
			os.Exit(1)
		}

		if err := xkcd.WriteJSONFile(file, comic); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to write JSON to file: %s\n", err)
			os.Exit(1)
		}

		if progress {
			fmt.Printf("Downloaded: Comic #%d\n", number)
		}
	}

	files, err := filepath.Glob(filepath.Join(cacheDir, "*.json"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not find comic cache files: %s\n", err)
		os.Exit(1)
	}

	var results []xkcd.Comic
	for _, file := range files {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			continue
		}

		if !strings.Contains(string(data), query) {
			continue
		}

		var comic xkcd.Comic
		if err := json.Unmarshal(data, &comic); err != nil {
			continue
		}

		results = append(results, comic)
	}

	if len(results) > 0 {
		fmt.Printf("Found %d matching comics: \n", len(results))
		for _, comic := range results {
			fmt.Printf(" - #%d %s: %s\n", comic.Number, comic.SafeTitle, comic.Img)
		}
	}
}
