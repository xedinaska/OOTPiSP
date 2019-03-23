#!make

BINARY = ./bin/app

lint:
	goimports -w $$(ls -d */ | grep -v vendor)
	golint $$(ls -d */ | grep -v vendor)
	gocyclo -over 10  $$(ls -d */ | grep -v vendor)

build:
	go build -o $(BINARY) cmd/main.go

run: build
	 ./$(BINARY)