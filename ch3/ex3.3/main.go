// Surface computes an SVG rendering of a 3-D function.
// Solution: Split out x,y,z calculation into one function.
//			 Split out corner into another function.
// 			 Use z value of point D to determine colour.
// 			 Apply colour as a fill on the polygon.
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30º)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30º), cos(30º)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7;' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(coords(i+1, j))
			bx, by := corner(coords(i, j))
			cx, cy := corner(coords(i, j+1))
			dx, dy, dz := coords(i+1, j+1)
			dx, dy = corner(dx, dy, dz)

			// Change colour depending on whether peak or valley.
			colour := "transparent"

			if dz > 0 {
				colour = "#ff0000"
			} else if dz < 0 {
				colour = "#0000ff"
			}

			fmt.Printf("<polygon points='%g,%g,%g,%g,%g,%g,%g,%g' fill='%s'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, colour)
		}
	}

	fmt.Printf("</svg>")
}

func coords(i, j int) (float64, float64, float64) {
	// Find point (x, y) at corner of cell (i, j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z
	z := f(x, y)

	return x, y, z
}

func corner(x, y, z float64) (float64, float64) {
	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx, sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale

	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // Distance from (0, 0)
	return math.Sin(r) / r
}
