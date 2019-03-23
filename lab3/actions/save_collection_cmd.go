package actions

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/xedinaska/OOTPiSP/lab3/model"
)

//SaveCollectionAction used to save collection to file
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
