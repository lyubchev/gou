package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"strconv"

	"github.com/kbinani/screenshot"
)

// Colors
const (
	Blue   = "#63c7ff"
	Yellow = "#f7ef39"
	Purple = "#bd75ff"
	Orange = "#ff8221"
	Green  = "#42f331"
	Pink   = "#ff82b5"
	Gray   = "#8c8a8c"
	White  = "#ffffff"
)

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

	// bounds := screenshot.GetDisplayBounds(0)
	bounds := image.Rect(x0, y0, x1, y1)

	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		panic(err)
	}

	file, err := os.Create("pesho.png")
	if err != nil {
		panic(err)
	}

	defer file.Close()
	png.Encode(file, img)

	fmt.Printf("#%d : %v \"%s\"\n", 0, bounds, "pesho.png")

	fmt.Println(bounds)
	// fmt.Println(img)
}
