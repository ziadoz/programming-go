// Fetch all fetches URLs in parallel and reports their times and sizes.
// 1.10 - Find a website that produces lot of data. Investigate caching.
//        Run twice and check if content is the same.
//        Modify so content is written to a file that can be examined.
package main

import (
    "fmt"
    "io"
    //"io/ioutil"
    "net/http"
    "os"
    "strings"
    "path/filepath"
    "time"
)

func main() {
    start := time.Now();
    ch    := make(chan string)

    for _, url := range os.Args[1:] {
        go fetch(url, ch) // Start a goroutine.
    }

    for range os.Args[1:] {
        fmt.Println(<-ch) // Receive from channel ch.
    }

    fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
    start := time.Now()

    resp, err := http.Get(url)
    if err != nil {
        ch <- fmt.Sprint(err) // Send to channel ch.
        return
    }

    nbytes := writeUrlToFile(url, resp.Body)
    resp.Body.Close()

    secs := time.Since(start).Seconds()
    ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}

func cleanUrl(url string) string {
    url = strings.TrimPrefix(url, "http://")
    url = strings.TrimPrefix(url, "https://")
    url = strings.TrimPrefix(url, "www.")
    url = strings.TrimSuffix(url, "/")
    url = strings.Replace(url, ".", "_", -1)
    return url
}

func writeUrlToFile(url string, body io.Reader) int64 {
    pwd, _   := os.Getwd()
    dir, _   := filepath.Abs(pwd)
    filename := filepath.Join(dir, cleanUrl(url) + ".txt")

    out, err := os.Create(filename)
    if err != nil {
        fmt.Fprintf(os.Stderr, "while creating file %s: %v\n", filename, err)
        os.Exit(1)
    }

    nbytes, err := io.Copy(out, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "while copying to %s: %v\n", filename, err)
        os.Exit(1)
    }

    out.Close()
    return nbytes
}
