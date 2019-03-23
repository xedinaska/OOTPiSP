package model

//Figure describes geometrical figure. In our app it should realize only 1 method - Draw()
type Figure interface {
	Name() string
	Configure() Figure
	Points() []*Point
}

//Figures - collection of model.Figure as a separate "class" (struct in Golang)
type Figures []Figure

//Point means "coordinate point"
type Point struct {
	X int
	Y int
}
