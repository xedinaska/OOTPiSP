package main

import (
	"context"
	"fmt"

	"flag"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/xedinaska/OOTPiSP/lab6/actions"
	"github.com/xedinaska/OOTPiSP/lab6/encoding"
	"github.com/xedinaska/OOTPiSP/lab6/model"
	"github.com/xedinaska/OOTPiSP/lab6/plugins"
)

var (
	registry []model.Shape
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{})
}

func main() {
	collection := &model.Shapes{
		Items: make(map[string]model.Shape),
	}

	shapes := []model.Shape{
		new(model.Circle).Init(),
		new(model.Triangle).Init(),
		new(model.Rectangle).Init(),
		new(model.Square).Init(),
		new(model.Trapezium).Init(),
		model.NewLineAdapter().Init(),
	}

	encoders := []encoding.Encoder{
		&encoding.Default{},
	}

	commands := []actions.Action{
		&actions.AddObjectAction{Collection: collection},
		&actions.EditObjectAction{Collection: collection},
		&actions.RemoveObjectAction{Collection: collection},
		&actions.SaveCollectionAction{Collection: collection},
		&actions.LoadCollectionAction{Collection: collection},
		&actions.ListCollectionAction{Collection: collection},
	}

	var flags plugins.Flag
	flag.Var(&flags, "plugins", "plugin name")
	flag.Parse()

	for _, name := range flags {
		p, err := plugins.Load(name)

		v, err := p.Lookup("Exported")
		if err != nil {
			logrus.Errorf("failed to load plugin `%s`: `%s`", name, err.Error())
			continue
		}

		if strings.Contains(name, "shapes") {
			eShape := v.(model.Shape).Init()
			shapes = append(shapes, eShape)
		} else if strings.Contains(name, "actions") {
			cmd := v.(actions.Action)
			cmd.SetCollection(collection)
			commands = append(commands, cmd)
		} else if strings.Contains(name, "encoders") {
			encoders = append(encoders, v.(encoding.Encoder))
		}
	}

	commands = append(commands, &actions.SetEncoderAction{Collection: collection, Encoders: encoders})
	commands = append(commands, &actions.AdapterExitAction{})

	//register available shape types
	registerShapes(shapes...)

	for {
		fmt.Printf("what do you want to do (enter number)?\n")
		for _, a := range commands {
			fmt.Printf("%d. %s\n", a.Value(), a.Name())
		}

		var actType int
		if _, err := fmt.Scanf("%d", &actType); err != nil {
			logrus.Fatalf("failed to scan user input: `%s`", err.Error())
		}

		var err error
	cmdLoop:
		for _, a := range commands {
			switch actType {
			case a.Value():
				err = a.Exec(context.Background())
				break cmdLoop
			}
		}
		if err != nil {
			logrus.Errorf("command execution failure: `%s`", err.Error())
		}
	}
}

func registerShapes(shapes ...model.Shape) {
	registry = make([]model.Shape, 0)

	for _, sh := range shapes {
		model.RegistryInstance().Add(sh)
	}
}
