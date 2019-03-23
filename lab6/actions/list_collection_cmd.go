package actions

import (
	"context"
	"fmt"

	"github.com/xedinaska/OOTPiSP/lab6/model"
)

type ListCollectionAction struct {
	Collection *model.Shapes
}

func (a *ListCollectionAction) Exec(ctx context.Context) error {
	for _, ss := range a.Collection.Items {
		fmt.Printf("item: [%v]\n", ss)
	}
	return nil
}

func (a *ListCollectionAction) Value() int {
	return ListCollection
}

func (a *ListCollectionAction) Name() string {
	return "list current collection"
}

func (a *ListCollectionAction) SetCollection(shapes *model.Shapes) {}
