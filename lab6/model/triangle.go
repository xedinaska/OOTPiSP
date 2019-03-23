package model

import (
	"github.com/satori/go.uuid"
	"github.com/xedinaska/OOTPiSP/lab6/encoding"
)

type Triangle struct {
	Polygon
	Id        uuid.UUID
	ShapeName string
	Data      string
	Encoder   encoding.Encoder
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
	t.Data = coords
}

func (t *Triangle) Coords() string {
	return t.Data
}
