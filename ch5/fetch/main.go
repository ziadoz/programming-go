package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func main() {
	for _, url := range os.Args[1:] {
		fmt.Println(fetch(url))
	}
}

// Fecth downloads the URL and returns the name and length of the local file.
func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}

	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}

	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}

	n, err = io.Copy(f, resp.Body)
	// Close file, but prefer error from copy, if any
	// Closing may return an error (e.g. on NFS), so need to check for that error also.
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}

	return local, n, err
}
