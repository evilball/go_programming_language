//Exercise 3.6: Supersampling is a technique to reduce the effect of pixelation by
//computing the color value at several points within each pixel and taking the average.
//The simplest method is to divide each pixel into four “subpixels.” Implement it.

//result - http://evilball.videodemons.com/img/supersampling.png
package main

import (
	"image"
	"image/color"
	"math/cmplx"
	"os"
	"image/png"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
	epsX                   = (xmax - xmin) / width
	epsY                   = (ymax - ymin) / height
)

func main() {
		img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			// Supersampling:
			var subpixels []color.Color

			shiftX := []float64{-epsX, epsX}
			shiftY := []float64{-epsY, epsY}
			for _, i := range shiftX {
				for _, j := range shiftY {
					z := complex(x+i, y+j)
					subpixels = append(subpixels, mandelbrot(z))
				}
			}
			img.Set(px, py, avg(subpixels))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func avg(colors []color.Color) color.Color {
	var r, g, b, a uint16
	colorsLen := uint32(len(colors))
	for _, c := range colors {
		_r, _g, _b, _a := c.RGBA()
		r += uint16(_r / colorsLen)
		g += uint16(_g / colorsLen)
		b += uint16(_b / colorsLen)
		a += uint16(_a / colorsLen)
	}
	return color.RGBA64{r, g, b, a}
}

func mandelbrot(z complex128) color.RGBA {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{0 + contrast*n, 0 + contrast*n, 255 - contrast*n, 255}
		}
	}
	return color.RGBA{0, 0, 0, 255}
}