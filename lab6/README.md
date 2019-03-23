# Lab work N 6

## Content info

Based on Lab5. Added 2 adapters:
- `AdapterExitAction` in `actions` package. Allows to use `Killer` struct (that terminates app) from `github.com/xedinaska/appkiller` project.
- `LineAdapter` in `model` package. Allows to use `Line` shape from lab2.

Patterns:
- `Command` pattern already used for collection actions execution.
- `Singleton` + `Repository` used for shapes registration in system (allows us to register new shapes without code change).

## Installation

- Install glide (https://github.com/Masterminds/glide)
- Run `glide install` to fetch vendors.
- Run `make plugin NAME=FOLDER_UNDER_plugins/src/_NAME/FILENAME` to build plugin code
- Run `make build` command - it'll compile to `bin` folder.
- Run ./bin/app --plugins FOLDER_UNDER_plugins/src/_NAME/FILENAME --plugins FOLDER_UNDER_plugins/src/_NAME/FILENAME_2

