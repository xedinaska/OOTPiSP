package model

import (
	"bytes"
	"encoding/gob"
	"errors"

	"fmt"

	"github.com/satori/go.uuid"
)

type Shape interface {
	Name() string
	Family() string
	Configure() (string, error)
	SetCoords(string)
	Init() Shape
	ID() uuid.UUID
}

type Shapes struct {
	Items map[string]Shape
}

func (s *Shapes) Len() int {
	return len(s.Items)
}

func (s *Shapes) Add(item Shape) {
	s.Items[item.ID().String()] = item
}

func (s *Shapes) Remove(id string) error {
	if _, ok := s.Items[id]; !ok {
		return errors.New("item with id `%s` isn't found")
	}
	delete(s.Items, id)
	return nil
}

func (s *Shapes) Serialize() (text []byte, err error) {
	b := &bytes.Buffer{}
	err = gob.NewEncoder(b).Encode(s)
	return b.Bytes(), err
}

func (s *Shapes) Unserialize(src []byte) error {
	newCol := new(Shapes)
	err := gob.NewDecoder(bytes.NewBuffer(src)).Decode(newCol)
	if err != nil {
		return err
	}
	s.Items = newCol.Items
	return nil
}

type Polygon struct{}

func (p Polygon) Family() string {
	return "Polygon"
}

func (p Polygon) Configure() (string, error) {
	fmt.Println("please provide points coords in format (x1;y1),(x2;y2),..")
	var coords string
	_, err := fmt.Scanf("%s", &coords)
	return coords, err
}

type Ellipsoid struct{}

func (e Ellipsoid) Family() string {
	return "Ellipsoid"
}

func (e Ellipsoid) Configure() (string, error) {
	fmt.Println("please provide center coords & Radius in format (x1;y1),R")
	var coords string
	_, err := fmt.Scanf("%s", &coords)
	return coords, err
}
