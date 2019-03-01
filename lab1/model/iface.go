package model

//Figure describes geometrical figure. In our app it should realize only 1 method - Draw()
type Figure interface {
	Draw()
}

//Figures - collection of model.Figure as a separate "class" (struct in Golang)
type Figures []Figure