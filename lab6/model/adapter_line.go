package model

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	lab2model "github.com/xedinaska/OOTPiSP/lab2/model"
)

func NewLineAdapter() *LineAdapter {
	return &LineAdapter{
		line: &lab2model.Line{
			Start: new(lab2model.Point),
			End:   new(lab2model.Point),
		},
	}
}

type LineAdapter struct {
	Polygon
	Id        uuid.UUID
	ShapeName string
	line      *lab2model.Line
}

func (la *LineAdapter) Name() string {
	return la.ShapeName
}

func (la *LineAdapter) Configure() (string, error) {
	la.line.Configure()
	return la.Coords(), nil
}

func (la *LineAdapter) Init() Shape {
	return &LineAdapter{
		Id:        uuid.NewV4(),
		ShapeName: "line",
		line:      la.line,
	}
}

func (la *LineAdapter) ID() uuid.UUID {
	return la.Id
}

func (la *LineAdapter) SetCoords(c string) {
	point := func(pointStr string) (x int, y int, err error) {
		coords := strings.Replace(
			strings.Replace(pointStr, ")", "", -1),
			"(",
			"",
			-1,
		)

		xy := strings.Split(coords, ";")
		x, err = strconv.Atoi(xy[0])
		if err != nil {
			return 0, 0, err
		}
		y, err = strconv.Atoi(xy[1])
		if err != nil {
			return 0, 0, err
		}

		return x, y, nil
	}

	parts := strings.Split(c, ",")
	if len(parts) != 2 {
		logrus.Errorf("invalid coords provided for line object")
		return
	}

	x1, y1, err := point(parts[0])
	if err != nil {
		logrus.Errorf("failed to parse first point: `%s`", err.Error())
		return
	}

	x2, y2, err := point(parts[1])
	if err != nil {
		logrus.Errorf("failed to parse second point: `%s`", err.Error())
		return
	}

	la.line.Start = &lab2model.Point{x1, y1}
	la.line.End = &lab2model.Point{x2, y2}
}

func (la *LineAdapter) Coords() string {
	coords := make([]string, 0)
	for _, p := range la.line.Points() {
		coords = append(coords, fmt.Sprintf("(%d;%d)", p.X, p.Y))
	}

	return strings.Join(coords, ",")
}
