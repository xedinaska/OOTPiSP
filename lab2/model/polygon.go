package model

import (
	"fmt"
	"strconv"
	"strings"
)

//Polygon is a just combination of few coords points
type Polygon struct {
	Nodes []*Point
}

//NewPolygon is smth like "constructor" for Polygon struct
func NewPolygon() *Polygon {
	return &Polygon{
		Nodes: make([]*Point, 0),
	}
}

//Configure asks user to fill in the line configuration (x0, y0, x1, y1, ...)
func (p *Polygon) Configure() Figure {
	fmt.Printf("Please provide comma-separated point coordinates, like x1,y1,x2,y2,x3,y3: ")

	var input string
	fmt.Scanf("%s", &input)

	values := strings.Split(input, ",")
	for i := 0; i < len(values)-1; i += 2 {
		x, _ := strconv.Atoi(values[i])
		y, _ := strconv.Atoi(values[i+1])
		p.Nodes = append(p.Nodes, &Point{X: x, Y: y})
	}

	return p
}

//Name returns a figure name (used to show during configuration process)
func (p *Polygon) Name() string {
	return "polygon"
}

//Points used to determine the points needed for rasterizing a rectangle.
func (p *Polygon) Points() []*Point {
	points := make([]*Point, 0)
	for i := 1; i < len(p.Nodes); i++ {
		l := &Line{
			Start: p.Nodes[i-1],
			End:   p.Nodes[i],
		}
		for _, pp := range l.Points() {
			points = append(points, pp)
		}
	}
	return points
}
