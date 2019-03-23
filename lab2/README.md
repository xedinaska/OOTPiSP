# Lab work N 2

## Content info

You can find all shapes inside ./model directory. Each shape is located in separate file but common interface is defined in `model/model.go` file.

Drawing functionality moved out of figures and described in `ui/ui.go` file (`Drawer` struct). Each shape returns slice of points and drawer draws them on canvas.
You need to open a browser on provided port to view result (:5000 by default).

## Installation

- Install glide (https://github.com/Masterminds/glide)
- Run `glide install` to fetch vendors.
- Run `make run` command - it'll compile and run app automatically.

