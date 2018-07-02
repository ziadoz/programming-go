package xkcd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const Url = "http://xkcd.com/%d/info.0.json"
const MaxUrl = "https://xkcd.com/info.0.json"

type Comic struct {
	Year       int    `json:"year,string"`
	Month      int    `json:"month,string"`
	Day        int    `json:"day,string"`
	Number     int    `json:"num"`
	Title      string `json:"title"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Img        string `json:"img"`
	Alt        string `json:"alt"`
}

func GetComic(number int) (*Comic, error) {
	return getComicFromURL(fmt.Sprintf(Url, number))
}

func GetLatestComicNumber() (int, error) {
	comic, err := getComicFromURL(MaxUrl)
	if err != nil {
		return 0, fmt.Errorf("Could not get max comic numer: %s", err)
	}

	return comic.Number, nil
}

func getComicFromURL(url string) (*Comic, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP query failed with status code: %d", response.StatusCode)
	}

	if contentType := response.Header.Get("Content-Type"); contentType != "application/json" {
		return nil, fmt.Errorf("HTTP query did not return JSON: %s", contentType)
	}

	var comic Comic
	if err := json.NewDecoder(response.Body).Decode(&comic); err != nil {
		return nil, err
	}

	return &comic, nil
}

func WriteJSONFile(file string, comic *Comic) error {
	json, err := json.MarshalIndent(comic, "", "    ")
	if err != nil {
		return fmt.Errorf("Failed to marshal comic %d to JSON: %s", comic.Number, err)
	}

	if err = ioutil.WriteFile(file, json, 0644); err != nil {
		return fmt.Errorf("Failed to write JSON to file: %s", err)
	}

	return nil
}
