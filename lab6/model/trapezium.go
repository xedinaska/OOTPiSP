package model

import (
	"github.com/satori/go.uuid"
	"github.com/xedinaska/OOTPiSP/lab6/encoding"
)

type Trapezium struct {
	Polygon
	Id        uuid.UUID
	ShapeName string
	Data      string
	Encoder   encoding.Encoder
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
	t.Data = coords
}

func (t *Trapezium) Coords() string {
	return t.Data
}
