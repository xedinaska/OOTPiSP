package actions

import (
	"context"
	"os"
	"syscall"

	"github.com/xedinaska/OOTPiSP/lab6/model"
	"github.com/xedinaska/appkiller/app"
)

type AdapterExitAction struct{}

func (AdapterExitAction) Exec(ctx context.Context) error {
	k := app.Killer{
		ExitChan: make(chan os.Signal, 0),
	}
	k.Listen()
	k.ExitChan <- syscall.SIGTERM
	return nil
}

func (AdapterExitAction) Value() int {
	return ExitApp
}

func (AdapterExitAction) Name() string {
	return "exit"
}

func (AdapterExitAction) SetCollection(shapes *model.Shapes) {}
