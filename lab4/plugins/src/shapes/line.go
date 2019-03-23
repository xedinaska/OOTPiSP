package main

import (
	"fmt"

	"github.com/satori/go.uuid"
	"github.com/xedinaska/OOTPiSP/lab4/model"
)

var Exported Line

type Line struct {
	Id        uuid.UUID
	ShapeName string
	Coords    string
}

func (s Line) Name() string {
	return s.ShapeName
}

func (s Line) Init() model.Shape {
	return &Line{
		Id:        uuid.NewV4(),
		ShapeName: "LINE",
	}
}

func (s Line) ID() uuid.UUID {
	return s.Id
}

func (s Line) SetCoords(coords string) {
	s.Coords = coords
}

func (s Line) Family() string {
	return "Polygon"
}

func (s Line) Configure() (string, error) {
	fmt.Println("please provide points coords in format (x1;y1),(x2;y2)")
	var coords string
	_, err := fmt.Scanf("%s", &coords)
	return coords, err
}
