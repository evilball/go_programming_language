//Exercise 3.4: Following the approach of the Lissajous example in Section 1.7,
//construct a web server that computes surfaces and writes SVG data to the client. Theserver must set the Content-Type header like this:
//w.Header().Set("Content-Type", "image/svg+xml")

// go run Ex3.4.go
// Example address: localhost:8000/server?width=1000&height=1200

package main

import (
	"math"
	"fmt"
	"net/http"
	"strconv"
	"log"
	"io"
)

const (
	defaultWidth, defaultHeight = 600, 320    // canvas size in pixels
	cells                       = 100         // number of grid cells
	xyrange                     = 30.0        //axis ranges
	angle                       = math.Pi / 6 // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)
var width, height int

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var err error

		widthParam := r.FormValue("width");
		if widthParam != "" {
			width, err = strconv.Atoi(widthParam);
			if err != nil {
				fmt.Printf("Error while parsing param %s to int: %v", widthParam, err)
			}
		} else {
			width = defaultWidth
		}

		heightParam := r.FormValue("height");
		if heightParam != "" {
			height, err = strconv.Atoi(heightParam);
			if err != nil {
				fmt.Printf("Error while parsing param %s to int: %v", heightParam, err)
			}
		} else {
			height = defaultWidth
		}

		w.Header().Set("Content-Type", "image/svg+xml")
		createSvg(width, height, w)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func createSvg(width, length int, out io.Writer) {
	fmt.Fprintf(out, "<svg	xmlns='http://www.w3.org/2000/svg' style='stroke: grey; strokewidth: 0.7' width='%d' height='%d'>", width, length)
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
			fmt.Fprintf(out, "<polygon style='fill: %s' points='%g, %g %g, %g %g, %g %g, %g'/>\n",
				color, ax, ay, bx, by, cx, cy, dx, dy)

		}
	}
	fmt.Fprintln(out, "</svg>")
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

	xyscale := float64(width) / 2 / xyrange
	zscale := float64(height) * 0.4
	sx := float64(width)/2 + (x-y)*cos30*xyscale
	sy := float64(height)/2 + (x+y)*sin30*xyscale - z*zscale
	belowZero := z < 0
	return sx, sy, belowZero
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	result := math.Sin(r) / r
	if isValid(result) {
		return result
	} else {
		return 0
	}
}
