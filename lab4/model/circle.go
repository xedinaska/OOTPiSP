package model

import "github.com/satori/go.uuid"

type Circle struct {
	Ellipsoid
	Id        uuid.UUID
	ShapeName string
	Coords    string
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
	c.Coords = coords
}
