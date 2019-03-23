package actions

import (
	"context"
	"fmt"

	"github.com/xedinaska/OOTPiSP/lab6/model"
)

type EditObjectAction struct {
	Collection *model.Shapes
}

func (a *EditObjectAction) Exec(ctx context.Context) error {
	listAction := &ListCollectionAction{
		Collection: a.Collection,
	}
	listAction.Exec(ctx)

	fmt.Printf("please select shape to edit (enter ID): ")

	var id string
	if _, err := fmt.Scanf("%s", &id); err != nil {
		return fmt.Errorf("input scan - %s", err.Error())
	}

	item, ok := a.Collection.Items[id]
	if !ok {
		return fmt.Errorf("object %s doesn't exists", id)
	}

	coords, err := item.Configure()
	if err != nil {
		return fmt.Errorf("shape configuration - %s", err.Error())
	}

	item.SetCoords(coords)
	return nil
}

func (a *EditObjectAction) Value() int {
	return EditObject
}

func (a *EditObjectAction) Name() string {
	return "edit object in collection"
}

func (a *EditObjectAction) SetCollection(shapes *model.Shapes) {}
