package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"

	"gopl.io/ch5/links"
)

var (
	host       string              // The original host name.
	dir        string              // Where the downloaded files will be stored.
	downloaded = map[string]bool{} // What's already been downloaded.
)

func main() {
	log.SetPrefix("downloadlinks: ")
	log.SetFlags(0)

	if len(os.Args) == 1 {
		log.Fatalf("missing URL arguments")
	}

	link := os.Args[1]
	host = getURLHost(link)
	dir = path.Join("files", host)

	if err := os.Mkdir(dir, 0744); os.IsNotExist(err) {
		log.Fatalf("could not mkdir: %s", err)
	}

	breadthFirst(crawl, []string{link})
}

func crawl(link string) []string {
	fmt.Println(link)
	err := download(link)
	if err != nil {
		log.Fatalf("failed to download URL: %s", err)
	}

	list, err := links.Extract(link)
	if err != nil {
		log.Print(err)
	}

	for _, link := range list {
		err := download(link)
		if err != nil {
			log.Fatalf("failed to download URL: %s", err)
		}
	}

	return list
}

func getURL(link string) *url.URL {
	url, err := url.Parse(link)
	if err != nil {
		log.Fatalf("invalid URL: %s", err)
	}

	return url
}

func getURLHost(link string) string {
	url := getURL(link)
	return url.Host
}

func getFilename(link string) string {
	url := getURL(link)
	filename := strings.Trim(url.Path, "/")

	if ext := filepath.Ext(url.Path); ext != "" {
		filename = strings.Replace(filename, ext, "", -1)
	}

	filename = strings.Replace(filename, "/", "-", -1)
	if filename == "" {
		filename = "index"
	}

	return strings.ToLower(filename) + ".html"
}

func download(link string) error {
	// Check matches original host.
	if getURLHost(link) != host {
		return nil
	}

	// Check not already downloaded.
	if downloaded[link] {
		return nil
	}

	response, err := http.Get(link)
	if err != nil {
		return fmt.Errorf("could not get: %s", err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid status code: %d", response.StatusCode)
	}

	file, err := os.Create(path.Join(dir, getFilename(link)))
	if err != nil {
		return fmt.Errorf("could not create file: %s", err)
	}

	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return fmt.Errorf("could not copy response to file: %s", err)
	}

	downloaded[link] = true
	return nil
}

// breadthFirst calls f for each item in the worklist/
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}
