package actions

import (
	"context"

	"errors"
	"fmt"

	"github.com/xedinaska/OOTPiSP/lab6/encoding"
	"github.com/xedinaska/OOTPiSP/lab6/model"
)

type SetEncoderAction struct {
	Collection *model.Shapes
	Encoders   []encoding.Encoder
}

func (a *SetEncoderAction) Exec(ctx context.Context) error {
	fmt.Println("please select encoder:")
	for i, e := range a.Encoders {
		fmt.Printf("%d. %s\n", i, e.Name())
	}

	var selection int
	if _, err := fmt.Scanf("%d", &selection); err != nil {
		return err
	}

	if selection >= len(a.Encoders) {
		return errors.New("invalid value")
	}

	a.Collection.SetEncoder(a.Encoders[selection])
	return nil
}

func (a *SetEncoderAction) Value() int {
	return SetEncoder
}

func (a *SetEncoderAction) Name() string {
	return "set encoder for collection"
}

func (a *SetEncoderAction) SetCollection(shapes *model.Shapes) {}
