package actions

import (
	"context"
	"fmt"

	"github.com/xedinaska/OOTPiSP/lab6/model"
)

type AddObjectAction struct {
	Collection *model.Shapes
}

func (a *AddObjectAction) Exec(ctx context.Context) error {
	fmt.Printf("please select shape to add (enter number): \n")
	for idx, s := range model.RegistryInstance().Items() {
		fmt.Printf("%d. %s\n", idx, s.Name())
	}

	var shapeType int
	if _, err := fmt.Scanf("%d", &shapeType); err != nil {
		return err
	}

	newObject := model.RegistryInstance().Get(shapeType)
	a.Collection.Add(newObject.Init())
	return nil
}

func (a *AddObjectAction) Value() int {
	return AddObject
}

func (a *AddObjectAction) Name() string {
	return "add new object to collection"
}

func (a *AddObjectAction) SetCollection(shapes *model.Shapes) {}
