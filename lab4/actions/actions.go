package actions

import (
	"context"

	"github.com/xedinaska/OOTPiSP/lab4/model"
)

const (
	AddObject = iota
	EditObject
	RemoveObject
	SaveCollection
	LoadCollection
	ListCollection
)

type Action interface {
	Exec(ctx context.Context) error
	Value() int
	Name() string
	SetCollection(shapes *model.Shapes)
}
