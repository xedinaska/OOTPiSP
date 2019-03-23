package model

import (
	"github.com/satori/go.uuid"
	"github.com/xedinaska/OOTPiSP/lab5/encoding"
)

type Square struct {
	Polygon
	Id        uuid.UUID
	ShapeName string
	Data      string
	Encoder   encoding.Encoder
}

func (s *Square) Name() string {
	return s.ShapeName
}

func (s *Square) Init() Shape {
	return &Square{
		Id:        uuid.NewV4(),
		ShapeName: "square",
	}
}

func (s *Square) ID() uuid.UUID {
	return s.Id
}

func (s *Square) SetCoords(coords string) {
	s.Data = coords
}

func (s *Square) Coords() string {
	return s.Data
}
