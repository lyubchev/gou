package main

import (
	"fmt"
	"image"
	"image/color"
	"os"
	"strconv"
	"time"

	"github.com/JamesHovious/w32"
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

var down = w32.INPUT{
	Type: 0,
	Mi: w32.MOUSEINPUT{
		DwFlags: w32.MOUSEEVENTF_LEFTDOWN,
	},
}

var up = w32.INPUT{
	Type: 0,
	Mi: w32.MOUSEINPUT{
		DwFlags: w32.MOUSEEVENTF_LEFTUP,
	},
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

	screenshotBounds := img.Bounds()
	width := screenshotBounds.Max.X
	height := screenshotBounds.Max.Y

	// BLM

	for i := 0; i < 1; i++ {
		var pouColor color.Color = color.Transparent
		setColor := true

		for y := screenshotBounds.Min.Y; y < height; y++ {
			for x := screenshotBounds.Min.X; x < width; x++ {

				pix := img.At(x, y)

				for _, c := range colors {
					if !setColor {
						break
					}

					if c == pix {
						pouColor = c
						setColor = false
					}
				}

				if !setColor && pouColor == pix {
					MoveClick(x0+x, y0+y)

					img, err = screenshot.CaptureRect(bounds)
					if err != nil {
						panic(err)
					}
					// x = screenshotBounds.Min.X
					// y = screenshotBounds.Min.Y

				}

				// if pix != pouColor && c == pix {

				// break
				// }

			}
		}
	}

	fmt.Println(bounds)
}

func MoveClick(x, y int) {
	w32.SetCursorPos(x, y)

	err := w32.SendInput([]w32.INPUT{down})
	if err != nil {
		panic(err)
	}

	time.Sleep(time.Second)

	err = w32.SendInput([]w32.INPUT{up})
	if err != nil {
		panic(err)
	}
}
