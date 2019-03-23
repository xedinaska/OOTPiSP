package actions

import (
	"context"

	"github.com/xedinaska/OOTPiSP/lab6/model"
)

const (
	AddObject = iota
	EditObject
	RemoveObject
	SaveCollection
	LoadCollection
	ListCollection
	SetEncoder
	ExitApp
)

type Action interface {
	Exec(ctx context.Context) error
	Value() int
	Name() string
	SetCollection(shapes *model.Shapes)
}
