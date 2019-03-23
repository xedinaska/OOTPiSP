package main

import (
	"github.com/xedinaska/OOTPiSP/lab1/model"
)

func main() {
	//allocate memory for slice of 6 geometrical figures
	figures := make(model.Figures, 6)

	//statically initialize them
	figures = model.Figures{
		&model.Line{
			Start: &model.Point{X: 10, Y: 10},
			End:   &model.Point{X: 20, Y: 20},
		},
		&model.Circle{
			Center: &model.Point{X: 10, Y: 10},
			R:      5,
		},
		&model.Rectangle{
			TLCorner: &model.Point{X: 10, Y: 10},
			BRCorner: &model.Point{X: 20, Y: 5},
		},
		&model.Square{
			TLCorner: &model.Point{X: 50, Y: 50},
			R:        10,
		},
		&model.Triangle{
			A: &model.Point{X: 50, Y: 50},
			B: &model.Point{X: 100, Y: 100},
			C: &model.Point{X: 150, Y: 50},
		},
		&model.Polygon{
			Points: []*model.Point{
				{X: 0, Y: 0},
				{X: 0, Y: 20},
				{X: 20, Y: 0},
			},
		},
	}

	//go through figures slice & draw them (Draw() method Overrided for each figure)
	for _, f := range figures {
		f.Draw()
	}
}
