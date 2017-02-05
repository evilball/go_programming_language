//Exercise 3.3: Color each polygon based on its height, so that the peaks are colored
//red (#ff0000) and the valleys blue (#0000ff).

// result - https://i.gyazo.com/7e45db3fdc27a0c0f56162ee845fe147.png

package main

import (
	"math"
	"fmt"
)
const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                //axis ranges
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("<svg	xmlns='http://www.w3.org/2000/svg' style='stroke: grey; strokewidth: 0.7' width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, belowZero := corner(i+1, j)
			bx, by, _ := corner(i, j)
			cx, cy, _ := corner(i, j+1)
			dx, dy, _ := corner(i+1, j+1)
			var color string
			if (belowZero) {
				color = "#ff0000"
			} else {
				color = "#0000ff"
			}
			fmt.Printf("<polygon style='fill: %s' points='%g, %g %g, %g %g, %g %g, %g'/>\n",
				color, ax, ay, bx, by, cx, cy, dx, dy)

		}
	}
	fmt.Println("</svg>")
}
func isValid(number float64) bool {
	if math.IsInf(number, 0) || math.IsNaN(number) {
		return false
	} else {
		return true
	}
}

func corner(i, j int) (float64, float64, bool) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := f(x, y)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	belowZero := z < 0
	return sx, sy, belowZero
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	result := math.Sin(r)/r
	if isValid(result) {
		return result
	} else {
		return 0
	}
}
