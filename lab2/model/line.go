package model

import (
	"fmt"
)

//Line describes a single line from point Start to point End
type Line struct {
	Start *Point
	End   *Point
}

//NewLine is smth like "constructor" for Line struct
func NewLine() *Line {
	return &Line{
		Start: new(Point),
		End:   new(Point),
	}
}

//Configure asks user to fill in the line configuration (x0, y0, x1, y1)
func (l *Line) Configure() Figure {
	fmt.Printf("Please provide comma-separated line point coordinates, like x1,y1,x2,y2: ")
	fmt.Scanf("%d,%d,%d,%d", &l.Start.X, &l.Start.Y, &l.End.X, &l.End.Y)
	return l
}

//Name returns a figure name (used to show during configuration process)
func (l *Line) Name() string {
	return "line"
}

//Points used to determine the points needed for rasterizing a line.
//Based on Digital differential analyzer algorithm: https://en.wikipedia.org/wiki/Line_drawing_algorithm
func (l *Line) Points() []*Point {
	points := make([]*Point, 0)

	start, end := *l.Start, *l.End
	if start.X > end.X {
		start, end = end, start
	}

	dx, dy := end.X-start.X, end.Y-start.Y

	//optimize for horizontal line
	if start.Y == end.Y {
		for x := start.X; dx != 0; dx-- {
			points = append(points, &Point{
				X: int(x),
				Y: int(l.Start.Y),
			})
			x++
		}
		return points
	}

	//optimize for vertical line
	if start.X == end.X {
		//because point is x-axis ordered, dx cannot be negative
		if dy < 0 {
			dy = -1 * dy
		}

		if start.Y > end.Y {
			start.Y, end.Y = end.Y, start.Y
		}

		for y := start.Y; dy != 0; dy-- {
			points = append(points, &Point{
				X: int(start.X),
				Y: int(y),
			})
			y++
		}
		return points
	}

	for x := start.X; x < end.X; x++ {
		y := start.Y + dy*(x-start.X)/dx
		points = append(points, &Point{
			X: x,
			Y: y,
		})
	}

	return points
}
