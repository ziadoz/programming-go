package omdb

// http://www.omdbapi.com/
import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const Url = "http://www.omdbapi.com/?apikey=%s&t=%s"

type APIKey string

type Movie struct {
	Title  string
	Poster string
}

func SearchPosterAPI(apikey APIKey, query string) (*Movie, error) {
	url := fmt.Sprintf(Url, apikey, url.QueryEscape(query))
	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("HTTP query failed: %s", err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP query failed with status code: %d", response.StatusCode)
	}

	if contentType := response.Header.Get("Content-Type"); !strings.Contains(contentType, "application/json") {
		return nil, fmt.Errorf("HTTP query did not return JSON: %s", contentType)
	}

	var movie Movie
	if err := json.NewDecoder(response.Body).Decode(&movie); err != nil {
		return nil, fmt.Errorf("Could not parse JSON: %s", err)
	}

	return &movie, nil
}

func (movie *Movie) DownloadPoster(filename string) error {
	response, err := http.Get(movie.Poster)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP query failed with status code: %d", response.StatusCode)
	}

	if contentType := response.Header.Get("Content-Type"); !strings.Contains(contentType, "image/") {
		return fmt.Errorf("HTTP query did not return an image: %s", err)
	}

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("Could not create image file: %s", err)
	}

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return fmt.Errorf("Could not write image file: %s", err)
	}

	return nil
}
