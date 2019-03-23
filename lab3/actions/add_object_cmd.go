package actions

import (
	"context"
	"fmt"

	"github.com/xedinaska/OOTPiSP/lab3/model"
)

//AddObjectAction used to add new element to collection
type AddObjectAction struct {
	Collection *model.Shapes
}

func (a *AddObjectAction) Exec(ctx context.Context) error {
	registry := ctx.Value("registry").([]model.Shape)

	fmt.Printf("please select shape to add (enter number): \n")
	for idx, s := range registry {
		fmt.Printf("%d. %s\n", idx, s.Name())
	}

	var shapeType int
	if _, err := fmt.Scanf("%d", &shapeType); err != nil {
		return err
	}

	newObject := registry[shapeType]
	a.Collection.Add(newObject.Init())
	return nil
}
