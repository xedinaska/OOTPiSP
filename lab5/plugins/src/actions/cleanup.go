package main

import (
	"context"

	"github.com/xedinaska/OOTPiSP/lab5/model"
)

var Exported CleanupCollectionAction

type CleanupCollectionAction struct {
	Collection *model.Shapes
}

func (a *CleanupCollectionAction) Exec(ctx context.Context) error {
	a.Collection.Items = make(map[string]model.Shape, 0)
	return nil
}

func (a *CleanupCollectionAction) Value() int {
	return 10
}

func (a *CleanupCollectionAction) Name() string {
	return "cleanup collection"
}

func (a *CleanupCollectionAction) SetCollection(shapes *model.Shapes) {
	a.Collection = shapes
}
