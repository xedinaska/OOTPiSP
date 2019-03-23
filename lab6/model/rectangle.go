package model

import (
	"github.com/satori/go.uuid"
	"github.com/xedinaska/OOTPiSP/lab6/encoding"
)

type Rectangle struct {
	Polygon
	Id        uuid.UUID
	ShapeName string
	Data      string
	Encoder   encoding.Encoder
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
	r.Data = coords
}

func (r *Rectangle) Coords() string {
	return r.Data
}
