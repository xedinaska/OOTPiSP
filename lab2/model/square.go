package model

import "fmt"

//Square is described by top left corner coords and side size (R)
type Square struct {
	TLCorner *Point
	R        int
}

//NewSquare is smth like "constructor" for Square struct
func NewSquare() *Square {
	return &Square{
		TLCorner: new(Point),
	}
}

//Configure asks user to fill in the line configuration (x0, y0, r)
func (s *Square) Configure() Figure {
	fmt.Printf("Please provide comma-separated point coordinates with radius, like x1,y1,r: ")
	fmt.Scanf("%d,%d,%d", &s.TLCorner.X, &s.TLCorner.Y, &s.R)
	return s
}

//Name returns a figure name (used to show during configuration process)
func (s *Square) Name() string {
	return "square"
}

//Points used to determine the points needed for rasterizing a square
func (s *Square) Points() []*Point {
	r := &Rectangle{
		TLCorner: s.TLCorner,
		BRCorner: &Point{
			X: s.TLCorner.X + s.R,
			Y: s.TLCorner.Y - s.R,
		},
	}
	return r.Points()
}
