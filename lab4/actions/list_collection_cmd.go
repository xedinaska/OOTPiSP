package actions

import (
	"context"
	"fmt"

	"github.com/xedinaska/OOTPiSP/lab4/model"
)

type ListCollectionAction struct {
	Collection *model.Shapes
}

func (a *ListCollectionAction) Exec(ctx context.Context) error {
	for _, ss := range a.Collection.Items {
		fmt.Printf("id: [%s], type: %s\n", ss.ID().String(), ss.Name())
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
