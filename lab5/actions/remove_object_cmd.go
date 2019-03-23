package actions

import (
	"context"
	"fmt"

	"github.com/xedinaska/OOTPiSP/lab5/model"
)

type RemoveObjectAction struct {
	Collection *model.Shapes
}

func (a *RemoveObjectAction) Exec(ctx context.Context) error {
	var shapeID string
	fmt.Printf("please select shape to remove (enter UUID): \n")
	if _, err := fmt.Scanf("%s", &shapeID); err != nil {
		return err
	}

	return a.Collection.Remove(shapeID)
}

func (a *RemoveObjectAction) Value() int {
	return RemoveObject
}

func (a *RemoveObjectAction) Name() string {
	return "remove object from collection"
}

func (a *RemoveObjectAction) SetCollection(shapes *model.Shapes) {}
