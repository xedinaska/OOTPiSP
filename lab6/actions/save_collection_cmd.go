package actions

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/xedinaska/OOTPiSP/lab6/model"
)

type SaveCollectionAction struct {
	Collection *model.Shapes
}

func (a *SaveCollectionAction) Exec(ctx context.Context) error {
	text, err := a.Collection.Serialize()
	if err != nil {
		return fmt.Errorf("collection marshal - %s", err.Error())
	}

	fmt.Printf("please provide name file name: ")
	var filename string
	if _, err := fmt.Scanf("%s", &filename); err != nil {
		return fmt.Errorf("input scan - %s", err.Error())
	}

	if err := ioutil.WriteFile(filename, text, 0777); err != nil {
		return fmt.Errorf("write data to file - %s", err.Error())
	}

	return nil
}

func (a *SaveCollectionAction) Value() int {
	return SaveCollection
}

func (a *SaveCollectionAction) Name() string {
	return "save collection to file"
}

func (a *SaveCollectionAction) SetCollection(shapes *model.Shapes) {}
