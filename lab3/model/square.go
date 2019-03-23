package model

import "github.com/satori/go.uuid"

type Square struct {
	Polygon
	Id        uuid.UUID
	ShapeName string
	Coords    string
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
	s.Coords = coords
}
