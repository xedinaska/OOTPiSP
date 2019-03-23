package ui

import (
	"image"
	"image/color"
	"image/png"
	"io"

	"github.com/xedinaska/OOTPiSP/lab2/model"
)

var defaultColor = color.RGBA{255, 0, 0, 255}

//Drawer is used to draw figure points collection on image
type Drawer struct {
	points []*model.Point
}

//NewDrawer is just a Drawer constructor
func NewDrawer(p []*model.Point) *Drawer {
	return &Drawer{
		points: p,
	}
}

//Draw creates image instance, mark points from drawer.points slice & writes result to forwards them to provided Writer
func (d *Drawer) Draw(w io.Writer) {
	i := image.NewRGBA(image.Rect(0, 0, 100, 100))

	for _, p := range d.points {
		i.Set(p.X, 100-p.Y, defaultColor)
	}

	png.Encode(w, i)
}
