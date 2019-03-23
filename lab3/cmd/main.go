package main

import (
	"context"
	"encoding/gob"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/xedinaska/OOTPiSP/lab3/actions"
	"github.com/xedinaska/OOTPiSP/lab3/model"
)

var (
	registry []model.Shape
)

var collectionActions = map[int]string{
	actions.AddObject:      "add new object to collection",
	actions.LoadCollection: "load collection from file",
	actions.ListCollection: "list current collection",
	actions.EditObject:     "edit object in collection",
	actions.RemoveObject:   "remove object from collection",
	actions.SaveCollection: "save collection to file",
}

func init() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{})
}

func main() {
	//register available shape types
	registerShapes(
		new(model.Circle).Init(),
		new(model.Triangle).Init(),
		new(model.Rectangle).Init(),
		new(model.Square).Init(),
		new(model.Trapezium).Init(),
	)

	collection := &model.Shapes{
		Items: make(map[string]model.Shape),
	}

	ctx := context.WithValue(context.Background(), "registry", registry)

	addObjectCommand := &actions.AddObjectAction{Collection: collection}
	editObjectCommand := &actions.EditObjectAction{Collection: collection}
	removeObjectCommand := &actions.RemoveObjectAction{Collection: collection}
	saveCollectionCommand := &actions.SaveCollectionAction{Collection: collection}
	loadCollectionCommand := &actions.LoadCollectionAction{Collection: collection}
	listCollectionCommand := &actions.ListCollectionAction{Collection: collection}

	for {
		fmt.Printf("what do you want to do (enter number)?\n")
		for act, q := range collectionActions {
			fmt.Printf("%d. %s\n", act, q)
		}

		var actType int
		if _, err := fmt.Scanf("%d", &actType); err != nil {
			logrus.Fatalf("failed to scan user input: `%s`", err.Error())
		}

		var err error
		switch actType {
		case actions.AddObject:
			err = addObjectCommand.Exec(ctx)
		case actions.EditObject:
			err = editObjectCommand.Exec(ctx)
		case actions.RemoveObject:
			err = removeObjectCommand.Exec(ctx)
		case actions.SaveCollection:
			err = saveCollectionCommand.Exec(ctx)
		case actions.LoadCollection:
			err = loadCollectionCommand.Exec(ctx)
		case actions.ListCollection:
			err = listCollectionCommand.Exec(ctx)
		default:
			logrus.Error("unsupported action")
			continue
		}

		if err != nil {
			logrus.Errorf("command execution failure: `%s`", err.Error())
		}
	}
}

func registerShapes(shapes ...model.Shape) {
	registry = make([]model.Shape, 0)

	for _, sh := range shapes {
		registry = append(registry, sh)
		gob.Register(sh)
	}
}
