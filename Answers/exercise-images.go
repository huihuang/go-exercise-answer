/* Exercise: Images */

package main

import (
	"code.google.com/p/go-tour/pic"
	"image"
	"image/color"
)

/*
type Image interface {
    ColorModel() color.Model
    Bounds() Rectangle
    At(x, y int) color.Color
}
*/

type Image struct{}

func (pic Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (pic Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 100, 100)
}

func (pic Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x*x+y*y), uint8(y*x), 255, 255}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
