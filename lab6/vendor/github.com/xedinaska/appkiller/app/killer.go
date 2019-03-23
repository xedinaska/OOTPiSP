package app

import "os"

type Killer struct {
	ExitChan chan os.Signal
}

func (k *Killer) Listen() {
	go func() {
		<-k.ExitChan
		os.Exit(0)
	}()
}