# Go parameters
BINARY = ./bin/app

all: build

lint:
	goimports -w $$(ls -d */ | grep -v vendor)
	golint $$(ls -d */ | grep -v vendor)
	gocyclo -over 10  $$(ls -d */ | grep -v vendor)

build:
	go build -o $(BINARY) cmd/main.go