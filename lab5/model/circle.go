package model

import (
	"github.com/satori/go.uuid"
	"github.com/xedinaska/OOTPiSP/lab5/encoding"
)

type Circle struct {
	Ellipsoid
	Id        uuid.UUID
	ShapeName string
	Data      string
	Encoder   encoding.Encoder
}

func (c *Circle) Name() string {
	return c.ShapeName
}

func (c *Circle) Init() Shape {
	return &Circle{
		Id:        uuid.NewV4(),
		ShapeName: "circle",
	}
}

func (c *Circle) ID() uuid.UUID {
	return c.Id
}

func (c *Circle) SetCoords(coords string) {
	c.Data = coords
}

func (c *Circle) Coords() string {
	return c.Data
}
