package server

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

//Serve is used to start a web server at provided port & render result to browser
func Serve(timeout time.Duration, port string, draw func(w io.Writer)) {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
			<!doctype html>
			<title>Example</title>
			<img src="/canvas.png">
		`))
	})

	router.HandleFunc("/canvas.png", func(w http.ResponseWriter, r *http.Request) {
		draw(w)
	})

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		Handler:      router,
		ReadTimeout:  timeout,
		WriteTimeout: timeout,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("error listening server: %s", err.Error())
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	logrus.Info("shutdown server ...")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logrus.Fatalf("error shutdown server: %s", err.Error())
	}

	logrus.Info("server exiting")
}
