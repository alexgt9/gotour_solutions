package main

import (
	"code.google.com/p/go-tour/pic"
	"image"
	"image/color"
)

type Image struct{}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 50, 50)
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) At(x, y int) color.Color {
	return color.RGBA{uint8((x * 2) % 255), uint8((x * y) % 255), uint8((x * 4) % 255), 255}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
