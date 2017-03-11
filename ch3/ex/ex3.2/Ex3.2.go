//Exercise 3.2: Experiment with visualizations of other functions from the math
//package. Can you produce an egg box, moguls, or a saddle?

// result - https://i.gyazo.com/8880705412457425fd548d3cc90a66bf.png

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
	fmt.Printf("<svg	xmlns='http://www.w3.org/2000/svg' style='stroke: grey; fill: white; strokewidth: 0.7' width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Printf("<polygon points='%g, %g %g, %g %g, %g %g, %g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
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

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := f(x, y)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

// Some random func
func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	b := math.Cos(r)/r
	if isValid(b) {
		a := math.Sin(r)/r
		if (b > math.J0(a)) {
			return -a
		} else {
			return -b
		}

	} else {
		return 0
	}
}
