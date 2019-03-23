package model

import (
	"fmt"
)

//Rectangle is described by top-left & bottom-right corner coords
type Rectangle struct {
	TLCorner *Point
	BRCorner *Point
}

//NewRectangle is smth like "constructor" for Rectangle struct
func NewRectangle() *Rectangle {
	return &Rectangle{
		TLCorner: new(Point),
		BRCorner: new(Point),
	}
}

//Configure asks user to fill in the line configuration (x0, y0, x1, y1)
func (r *Rectangle) Configure() Figure {
	fmt.Printf("Please provide comma-separated point coordinates, like x1,y1,x2,y2: ")
	fmt.Scanf("%d,%d,%d,%d", &r.TLCorner.X, &r.TLCorner.Y, &r.BRCorner.X, &r.BRCorner.Y)
	return r
}

//Name returns a figure name (used to show during configuration process)
func (r *Rectangle) Name() string {
	return "rectangle"
}

//Points used to determine the points needed for rasterizing a rectangle.
func (r *Rectangle) Points() []*Point {
	points := make([]*Point, 0)

	l1 := &Line{
		Start: r.TLCorner,
		End:   &Point{X: r.BRCorner.X, Y: r.TLCorner.Y},
	}
	l2 := &Line{
		Start: l1.End,
		End:   r.BRCorner,
	}
	l3 := &Line{
		Start: r.BRCorner,
		End:   &Point{X: r.TLCorner.X, Y: r.BRCorner.Y},
	}
	l4 := &Line{
		Start: &Point{X: r.TLCorner.X, Y: r.BRCorner.Y},
		End:   r.TLCorner,
	}

	for _, l := range []*Line{l1, l2, l3, l4} {
		for _, p := range l.Points() {
			points = append(points, p)
		}
	}

	return points
}
