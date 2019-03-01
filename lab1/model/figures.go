package model

import (
	"fmt"
	"log"
	"strings"
)

//Point means "coordinate point"
type Point struct {
	X int
	Y int
}

//Line - just a single line A - B
type Line struct {
	Start *Point
	End   *Point
}

func (l *Line) Draw() {
	log.Printf("Line([%d, %d], [%d, %d])", l.Start.X, l.Start.Y, l.End.X, l.End.Y)
}

//Circle with center in "Center" point and it has radius R
type Circle struct {
	Center *Point
	R int
}

func (c *Circle) Draw() {
	log.Printf("Circle([%d, %d], %d)\n", c.Center.X, c.Center.Y, c.R)
}

//Rectangle is described by top-left & bottom-right corner coords
type Rectangle struct {
	TLCorner *Point
	BRCorner *Point
}

func (r *Rectangle) Draw() {
	log.Printf("Rectangle([%d, %d], [%d, %d])\n", r.TLCorner.X, r.TLCorner.Y, r.BRCorner.X, r.BRCorner.Y)
}

//Ellipse is described by center coords and radius in Ox & Oy
type Ellipse struct {
	Center *Point
	Xr int
	Yr int
}

func (e *Ellipse) Draw() {
	log.Printf("Ellipse([%d, %d], %d/X, %d/Y)", e.Center.X, e.Center.Y, e.Xr, e.Yr)
}

//Square is described by top left corner coords and side size (R)
type Square struct {
	TLCorner *Point
	R int
}

func (s *Square) Draw() {
	log.Printf("Square([%d, %d], %d)", s.TLCorner.X, s.TLCorner.Y, s.R)
}

//Polygon is a just combination of few coords points
type Polygon struct {
	Points []*Point
}

func (p *Polygon) Draw() {
	coords := func(points []*Point) []string {
		res := make([]string, len(points))
		for i, ps := range points {
			res[i] = fmt.Sprintf("[%d, %d]", ps.X, ps.Y)
		}
		return res
	}(p.Points)

	log.Printf("Polygon(%v)", strings.Join(coords, ","))
}