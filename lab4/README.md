# Lab work N 4

## Content info

Based on Lab3. Added new shape: `plugins/src/shapes/line.go`. Added new action: collection cleanup (`plugins/src/actions/cleanup.go`)

## Installation

- Install glide (https://github.com/Masterminds/glide)
- Run `glide install` to fetch vendors.
- Run `make plugin NAME=FOLDER_UNDER_plugins/src/_NAME/FILENAME` to build plugin code
- Run `make build` command - it'll compile to `bin` folder.
- Run ./bin/app --plugins FOLDER_UNDER_plugins/src/_NAME/FILENAME --plugins FOLDER_UNDER_plugins/src/_NAME/FILENAME_2

