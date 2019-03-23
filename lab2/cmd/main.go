package main

import (
	"fmt"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/xedinaska/OOTPiSP/lab2/model"
	"github.com/xedinaska/OOTPiSP/lab2/server"
	"github.com/xedinaska/OOTPiSP/lab2/ui"
)

const timeout = 5 * time.Second
const defaultHTTPPort = "3000"

//init is used to initialize application logger
func init() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {
	figures := model.Figures{
		model.NewCircle(),
		model.NewLine(),
		model.NewRectangle(),
		model.NewTriangle(),
		model.NewSquare(),
		model.NewPolygon(),
	}

	for i, f := range figures {
		fmt.Printf("Please enter `%d` to draw a %s\n", i, f.Name())
	}

	var pos int
	if _, err := fmt.Scanf("%d", &pos); err != nil {
		logrus.Fatalf("failed to scan input: `%s`", err.Error())
	}

	if pos < 0 || pos > len(figures)-1 {
		logrus.Fatalf("you should enter number between 0 and %d", len(figures)-1)
	}

	figure := figures[pos].Configure()

	port := os.Getenv("http_port")
	if port == "" {
		port = defaultHTTPPort
	}

	fmt.Printf("Generated image should be available here: http://127.0.0.1:%s\n", port)

	server.Serve(timeout, port, ui.NewDrawer(figure.Points()).Draw)
}
