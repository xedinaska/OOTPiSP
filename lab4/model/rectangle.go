package model

import "github.com/satori/go.uuid"

type Rectangle struct {
	Polygon
	Id        uuid.UUID
	ShapeName string
	Coords    string
}

func (r *Rectangle) Name() string {
	return r.ShapeName
}

func (r *Rectangle) Init() Shape {
	return &Rectangle{
		Id:        uuid.NewV4(),
		ShapeName: "rectangle",
	}
}

func (r *Rectangle) ID() uuid.UUID {
	return r.Id
}

func (r *Rectangle) SetCoords(coords string) {
	r.Coords = coords
}
