package model

import "github.com/satori/go.uuid"

type Triangle struct {
	Polygon
	Id        uuid.UUID
	ShapeName string
	Coords    string
}

func (t *Triangle) Name() string {
	return t.ShapeName
}

func (t *Triangle) Init() Shape {
	return &Triangle{
		Id:        uuid.NewV4(),
		ShapeName: "triangle",
	}
}

func (t *Triangle) ID() uuid.UUID {
	return t.Id
}

func (t *Triangle) SetCoords(coords string) {
	t.Coords = coords
}
