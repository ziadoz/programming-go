// Server4 renders a lissajous GIF to a webpage.
// 1.12 - Setup web sever to rendrer lissajous images and allow query string to
//        adjust the cycles, size etc., and use the strconv.Atoi function to
//        convert a string to an integer.
package main

import (
    "image"
    "image/color"
    "image/gif"
    "io"
    "log"
    "math"
    "math/rand"
    "net/http"
    "strconv"
    "strings"
    "time"
)

var palette = []color.Color{color.White, color.Black}

const (
    whiteIndex = 0 // First colour in palette.
    blackIndex = 1 // Next colour in palette.
)

func main() {
    rand.Seed(time.Now().UTC().UnixNano())

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if err := r.ParseForm(); err != nil {
            log.Print(err)
        }

        cycles := 5 // Number of complete x oscillator revolutions.
        if val, ok := r.Form["cycles"]; ok {
            cycles, _ = strconv.Atoi(strings.Join(val, ""))
        }

        size := 100 // Image canvas covers [-size..+size]
        if val, ok := r.Form["size"]; ok {
            size, _ = strconv.Atoi(strings.Join(val, ""))
        }

        lissajous(w, cycles, size)
    })

    log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajous(out io.Writer, cycles int, size int) {
    const (
        res     = 0.001 // Angular resolution.
        nFrames = 64    // Number of animation frames.
        delay   = 8     // Delay between frames in 10ms units.
    )

    freq  := rand.Float64() * 3.0 // Relative frequency of y oscillator.
    anim  := gif.GIF{LoopCount: nFrames}
    phase := 0.0 // Phase difference.

    for i := 0; i < nFrames; i++ {
        rect := image.Rect(0, 0, 2*size+1, 2*size+1)
        img  := image.NewPaletted(rect, palette)

        for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
            x := math.Sin(t)
            y := math.Sin(t*freq + phase)

            img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), blackIndex)
        }

        phase += 0.1
        anim.Delay = append(anim.Delay, delay)
        anim.Image = append(anim.Image, img)
    }

    gif.EncodeAll(out, &anim) // Note: Ignoring encoding errors.
}
