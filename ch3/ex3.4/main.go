// Surface computes an SVG rendering of a 3-D function.
package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
)

const angle = math.Pi / 6                           // angle of x, y axes (=30º)
var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30º), cos(30º)

func main() {
	host := "localhost:8000"

	fmt.Println("Running web server on http://" + host)
	fmt.Println("Press Ctrl+C to exit")

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/favicon.ico", faviconHandler)

	log.Fatal(http.ListenAndServe(host, nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatal("Failed to parse query string")
	}

	width, err := strconv.Atoi(queryValueOrDefault(r, "width", "600"))
	if err != nil {
		log.Fatalf("Could not parse width querystring parameter: %s", err)
	}

	height, err := strconv.Atoi(queryValueOrDefault(r, "height", "320"))
	if err != nil {
		log.Fatalf("Could not parse height querystring parameter: %s", err)
	}

	colour := "#" + queryValueOrDefault(r, "colour", "white")

	svg := surface(width, height, colour)
	w.Header().Set("Content-Type", "image/svg+xml")
	w.Write([]byte(svg))
}

// Need to handle this as every browser sends a request for it automatically.
// If we don't handle it, the indexHandler gets it and the application dies.
func faviconHandler(w http.ResponseWriter, r *http.Request) {
	// Do nothing.
}

func queryValueOrDefault(r *http.Request, field, defaultValue string) string {
	value := r.URL.Query().Get(field)
	if value == "" {
		return defaultValue
	}

	return value
}

func surface(width, height int, colour string) string {
	cells := 100                            // number of grid cells
	xyrange := 30.0                         // axis ranges (-xyrange..+xyrange)
	xyscale := float64(width) / 2 / xyrange // pixels per x or y unit
	zscale := float64(height) * 0.4         // pixels per z unit

	svg := fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: %s; stroke-width: 0.7;' "+
		"width='%d' height='%d'>", colour, width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, width, height, cells, xyrange, xyscale, zscale)
			bx, by := corner(i, j, width, height, cells, xyrange, xyscale, zscale)
			cx, cy := corner(i, j+1, width, height, cells, xyrange, xyscale, zscale)
			dx, dy := corner(i+1, j+1, width, height, cells, xyrange, xyscale, zscale)

			svg += fmt.Sprintf("<polygon points='%g,%g,%g,%g,%g,%g,%g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}

	svg += fmt.Sprintf("</svg>")
	return svg
}

func corner(i, j, width, height, cells int, xyrange, xyscale, zscale float64) (float64, float64) {
	// Find point (x, y) at corner of cell (i, j)
	x := xyrange * (float64(i)/float64(cells) - 0.5)
	y := xyrange * (float64(j)/float64(cells) - 0.5)

	// Compute surface height z
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx, sy).
	sx := float64(width)/2 + (x-y)*cos30*xyscale
	sy := float64(height)/2 + (x+y)*sin30*xyscale - z*zscale

	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // Distance from (0, 0)
	return math.Sin(r) / r
}
