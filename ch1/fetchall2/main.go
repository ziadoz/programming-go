// Fetch all fetches URLs in parallel and reports their times and sizes.
// 1.10 - Find a website that produces lot of data. Investigate caching.
//        Run twice and check if content is the same.
//        Modify so content is written to a file that can be examined.
// 1.11 - Try with lots of arguments, such as the Alexa top 1 million websites
//        and see how the programme behaves.
//
// Alexa Top 1m: http://s3.amazonaws.com/alexa-static/top-1m.csv.zip
// Usage: go run main.go `< top100.csv`
package main

import (
    "fmt"
    "io"
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
        go fetch(fixUrl(url), ch) // Start a goroutine.
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

func fixUrl(url string) string {
    if ! strings.HasPrefix(url, "http://") || ! strings.HasPrefix(url, "https://") {
        return "http://" + url
    }

    return url
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
    filename := filepath.Join(dir, "content", cleanUrl(url) + ".txt")

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
