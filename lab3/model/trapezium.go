package model

import "github.com/satori/go.uuid"

type Trapezium struct {
	Polygon
	Id        uuid.UUID
	ShapeName string
	Coords    string
}

func (t *Trapezium) Name() string {
	return t.ShapeName
}

func (t *Trapezium) Init() Shape {
	return &Trapezium{
		Id:        uuid.NewV4(),
		ShapeName: "trapezium",
	}
}

func (t *Trapezium) ID() uuid.UUID {
	return t.Id
}

func (t *Trapezium) SetCoords(coords string) {
	t.Coords = coords
}
