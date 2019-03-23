package model

import (
	"fmt"
)

//Circle with center in "Center" point and it has radius R
type Circle struct {
	Center *Point
	R      int
}

//NewCircle is smth like "constructor" for Circle struct
func NewCircle() *Circle {
	return &Circle{
		Center: new(Point),
		R:      0,
	}
}

//Configure asks user to fill in the circle configuration (x0, y0, R)
func (c *Circle) Configure() Figure {
	fmt.Printf("Please provide comma-separated center (x, y) coords and radius size, like X,Y,R: ")
	fmt.Scanf("%d,%d,%d", &c.Center.X, &c.Center.Y, &c.R)
	return c
}

//Name returns a figure name (used to show during configuration process)
func (c *Circle) Name() string {
	return "circle"
}

//Points used to determine the points needed for rasterizing a circle.
//Based on Midpoint circle algorithm: https://en.wikipedia.org/wiki/Midpoint_circle_algorithm
func (c *Circle) Points() []*Point {
	x, y, dx, dy := c.R-1, 0, 1, 1
	e := dx - (c.R * 2)

	points := make([]*Point, 0)

	for x > y {
		points = append(points, &Point{
			X: c.Center.X + x,
			Y: c.Center.Y + y,
		})
		points = append(points, &Point{
			X: c.Center.X + y,
			Y: c.Center.Y + x,
		})
		points = append(points, &Point{
			X: c.Center.X - y,
			Y: c.Center.Y + x,
		})
		points = append(points, &Point{
			X: c.Center.X - x,
			Y: c.Center.Y + y,
		})
		points = append(points, &Point{
			X: c.Center.X - x,
			Y: c.Center.Y - y,
		})
		points = append(points, &Point{
			X: c.Center.X - y,
			Y: c.Center.Y - x,
		})
		points = append(points, &Point{
			X: c.Center.X + y,
			Y: c.Center.Y - x,
		})
		points = append(points, &Point{
			X: c.Center.X + x,
			Y: c.Center.Y - y,
		})

		if e <= 0 {
			y++
			e += dy
			dy += 2
		}

		if e > 0 {
			x--
			dx += 2
			e += dx - (c.R * 2)
		}
	}

	return points
}
