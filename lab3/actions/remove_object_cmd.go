package actions

import (
	"context"
	"fmt"

	"github.com/xedinaska/OOTPiSP/lab3/model"
)

//RemoveObjectAction used to remove object from collection
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
