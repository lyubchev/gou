package main

import (
	"fmt"
	"image"
	"image/color"
	"os"
	"strconv"

	"github.com/kbinani/screenshot"
)

// Pou Colors

var colors = []color.Color{
	color.RGBA{99, 199, 255, 255},
	color.RGBA{247, 239, 57, 255},
	color.RGBA{189, 117, 255, 255},
	color.RGBA{255, 130, 33, 255},
	color.RGBA{66, 243, 49, 255},
	color.RGBA{255, 130, 181, 255},
	color.RGBA{140, 138, 140, 255},
	color.RGBA{255, 255, 255, 255},
}

func main() {
	x0, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}
	y0, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic(err)
	}
	w, err := strconv.Atoi(os.Args[3])
	if err != nil {
		panic(err)
	}
	h, err := strconv.Atoi(os.Args[4])
	if err != nil {
		panic(err)
	}

	x1 := w + x0
	y1 := h + y0

	bounds := image.Rect(x0, y0, x1, y1)
	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		panic(err)
	}

	width := bounds.Max.X
	height := bounds.Max.Y

	for x := bounds.Min.X; x < width; x++ {
		for y := bounds.Min.Y; y < height; y++ {
			r, g, b, _ := img.At(x, y).RGBA()

		}
	}

	fmt.Println(bounds)
	// fmt.Println(img)
}
