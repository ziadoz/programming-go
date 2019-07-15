// Exercise 8.11 - Make a function that returns the first URL in a list to download.
package main

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://www.bbc.co.uk",
		"http://www.theverge.com",
		"https://www.theguardian.com",
		"https://reddit.com",
		"http://www.thelocal.com",
		"http://www.omgubuntu.co.uk",
	}

	result := getFastestUrl(links)
	fmt.Printf("got fastest result: %s\n", result)
}

func getFastestUrl(urls []string) string {
	results := make(chan string)
	ctx, cancel := context.WithCancel(context.Background())

	for _, url := range urls {
		go fetchUrl(ctx, url, results)
	}

	result := <-results
	cancel()

	return result
}

func fetchUrl(ctx context.Context, url string, results chan<- string) {
	start := time.Now()
	randomSleep() // This keeps the results interesting.

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("could not create new request: %s\n", err)
	}

	request = request.WithContext(ctx)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatalf("could not fetch url: %s\n", err)
	}
	defer response.Body.Close()

	nbytes, err := io.Copy(ioutil.Discard, response.Body)
	if err != nil {
		log.Fatalf("could not copy response body: %s\n", err)
	}

	results <- fmt.Sprintf("%s %.2fs %7d", url, time.Since(start).Seconds(), nbytes)
}

func randomSleep() {
	delay := rand.Intn(1000)
	time.Sleep(time.Millisecond * time.Duration(delay))
}
