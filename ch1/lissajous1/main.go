// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
    "image"
    "image/color"
    "image/gif"
    "io"
    "math"
    "math/rand"
    "os"
    "time"
)

var palette = []color.Color{color.White, color.Black}

const (
    whiteIndex = 0 // First colour in palette.
    blackIndex = 1 // Next colour in palette.
)

func main() {
    rand.Seed(time.Now().UTC().UnixNano())
    lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
    const (
        cycles  = 5     // Number of complete x oscillator revolutions.
        res     = 0.001 // Angular resolution.
        size    = 100   // Image canvas covers [-size..+size]
        nFrames = 64    // Number of animation frames.
        delay   = 8     // Delay between frames in 10ms units.
    )

    freq  := rand.Float64() * 3.0 // Relative frequency of y oscillator.
    anim  := gif.GIF{LoopCount: nFrames}
    phase := 0.0 // Phase difference.

    for i := 0; i < nFrames; i++ {
        rect := image.Rect(0, 0, 2*size+1, 2*size+1)
        img  := image.NewPaletted(rect, palette)

        for t := 0.0; t < cycles*2*math.Pi; t += res {
            x := math.Sin(t)
            y := math.Sin(t*freq + phase)

            img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
        }

        phase += 0.1
        anim.Delay = append(anim.Delay, delay)
        anim.Image = append(anim.Image, img)
    }

    gif.EncodeAll(out, &anim) // Note: Ignoring encoding errors.
}
