// Fetch prints the content found at each specified URL.
// 1.7 - Use io.Copy(dst, src) to write the output directly to the stdout.
// 1.8 Modify so that URLs are automatically prefixed with http://.
// 1.9 Modify so that the HTTP status code is printed out.
package main

import (
    "fmt"
    "io"
    "net/http"
    "os"
    "strings"
)

func main() {
    for _, url := range os.Args[1:] {
        if ! strings.HasPrefix(url, "http://") || ! strings.HasPrefix(url, "https://") {
            url = "http://" + url
        }

        resp, httperr := http.Get(url)
        if httperr != nil {
            fmt.Fprintf(os.Stderr, "fetch: %v\n", httperr)
            os.Exit(1)
        }

        fmt.Println("HTTP Status Code: ", resp.Status)

        _, cperr := io.Copy(os.Stdout, resp.Body)
        resp.Body.Close()
        if cperr != nil {
            fmt.Fprintf(os.Stderr, "fetch: %v\n", cperr)
            os.Exit(1)
        }
    }
}
