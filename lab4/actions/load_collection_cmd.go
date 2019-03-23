package actions

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/xedinaska/OOTPiSP/lab4/model"
)

type LoadCollectionAction struct {
	Collection *model.Shapes
}

func (a *LoadCollectionAction) Exec(ctx context.Context) error {
	fmt.Printf("please provide name file name: ")

	var filename string
	if _, err := fmt.Scanf("%s", &filename); err != nil {
		return fmt.Errorf("input scan - %s", err.Error())
	}

	f, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("open collection dump - %s", err.Error())
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return fmt.Errorf("read collection dump - %s", err.Error())
	}

	if err := a.Collection.Unserialize(b); err != nil {
		return fmt.Errorf("collection unmarshal - %s", err.Error())
	}

	return nil
}

func (a *LoadCollectionAction) Value() int {
	return LoadCollection
}

func (a *LoadCollectionAction) Name() string {
	return "load collection from file"
}

func (a *LoadCollectionAction) SetCollection(shapes *model.Shapes) {}
