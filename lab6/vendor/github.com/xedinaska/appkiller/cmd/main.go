package main

import (
	"github.com/sirupsen/logrus"
	"time"
	"github.com/xedinaska/appkiller/app"
	"os"
	"syscall"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.Debugf("starting my awesome app")

	k := &app.Killer{
		ExitChan: make(chan os.Signal, 1),
	}
	k.Listen()

	i := 0
	for {
		select {
		default:
			i++
			time.Sleep(time.Second*1)
			logrus.Debugf("for{}")
			if i > 5 {
				k.ExitChan<-os.Signal(syscall.SIGTERM)
			}
		}
	}
}