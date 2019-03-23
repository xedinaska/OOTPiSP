package actions

import (
	"context"
	"fmt"

	"github.com/xedinaska/OOTPiSP/lab3/model"
)

// ListCollectionAction draws collection items
type ListCollectionAction struct {
	Collection *model.Shapes
}

func (a *ListCollectionAction) Exec(ctx context.Context) error {
	for _, ss := range a.Collection.Items {
		fmt.Printf("id: [%s], type: %s\n", ss.ID().String(), ss.Name())
	}
	return nil
}
