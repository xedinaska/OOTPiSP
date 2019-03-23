package model

import (
	"fmt"
)

//Triangle is described by thee triangle points (A,B,C)
type Triangle struct {
	A *Point
	B *Point
	C *Point
}

//NewTriangle is smth like "constructor" for Triangle struct
func NewTriangle() *Triangle {
	return &Triangle{
		A: new(Point),
		B: new(Point),
		C: new(Point),
	}
}

//Configure asks user to fill in the line configuration (x0, y0, x1, y1, ...)
func (t *Triangle) Configure() Figure {
	fmt.Printf("Please provide comma-separated point coordinates, like x1,y1,x2,y2,x3,y3: ")
	fmt.Scanf("%d,%d,%d,%d,%d,%d", &t.A.X, &t.A.Y, &t.B.X, &t.B.Y, &t.C.X, &t.C.Y)
	return t
}

//Name returns a figure name (used to show during configuration process)
func (t *Triangle) Name() string {
	return "triangle"
}

//Points used to determine the points needed for rasterizing a triangle.
func (t *Triangle) Points() []*Point {
	points := make([]*Point, 0)

	l1 := &Line{
		Start: t.A,
		End:   t.B,
	}
	l2 := &Line{
		Start: t.B,
		End:   t.C,
	}
	l3 := &Line{
		Start: t.C,
		End:   t.A,
	}

	for _, l := range []*Line{l1, l2, l3} {
		for _, p := range l.Points() {
			points = append(points, p)
		}
	}

	return points
}
