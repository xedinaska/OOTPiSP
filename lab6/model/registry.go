package model

import "encoding/gob"

var instance *registry

type registry struct {
	shapes []Shape
}

func RegistryInstance() *registry {
	if instance == nil {
		instance = &registry{}
	}
	return instance
}

func (r *registry) Add(shape Shape) {
	r.shapes = append(r.shapes, shape)
	gob.Register(shape)
}

func (r *registry) Items() []Shape {
	return r.shapes
}

func (r *registry) Get(key int) Shape {
	return r.shapes[key]
}
