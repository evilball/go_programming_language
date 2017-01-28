//Exercise 1.12: Modify the Lissajous server to read parameter values from the URL.
//For example, you might arrange it so that a URL like
//http://localhost:8000/?cycles=20 sets the number of cycles to 20
//instead of the default 5. Use the strconv.Atoi function to convert the string
//parameter into an integer. You can see its documentation with go doc
//strconv.Atoi.

package main

import (
	"net/http"
	"fmt"
	"io"
	"image/gif"
	"image"
	"math"
	"math/rand"
	"sync"
	"image/color"
	"strconv"
	"log"
)

var mu sync.Mutex
var count int

var palette = []color.Color{color.Black,                        // Black = 0
                            color.White,                        // White = 1
                            color.RGBA{0xff, 0x00, 0x00, 0x01}, // Red = 2
                            color.RGBA{0x00, 0xff, 0x00, 0x01}, // Green = 3
                            color.RGBA{0x00, 0x00, 0xff, 0x01}} // Blue = 4

const (
	sizeOfPalette = 5
	defaultCycles = 5
	res     = 0.001 // angular resolution
	size    = 100   //image canvas covers [-size..+size]
	nframes = 64    //number of animation frames
	delay   = 8     // delay between frames in 10ms units
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if cyclesParam := r.FormValue("cycles"); cyclesParam != "" {
			if cycles, err := strconv.Atoi(cyclesParam); err == nil {
				lissajous(cycles, w)
			} else {
				fmt.Printf("Error while parsing param %s to int: %v", cyclesParam, err)
			}
		} else {
			lissajous(defaultCycles, w)
		}
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajous(cycles int, out io.Writer) {
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 //phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			colorIndex := uint8(i % sizeOfPalette)
			img.SetColorIndex(size+int(x*size + 0.5), size+int(y*size + 0.5), colorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
