#!make

BINARY = ./bin/app

lint:
	goimports -w $$(ls -d */ | grep -v vendor)
	golint $$(ls -d */ | grep -v vendor)
	gocyclo -over 10  $$(ls -d */ | grep -v vendor)

build:
	go build -o $(BINARY) cmd/main.go

plugin:
	go build -buildmode=plugin -o ./plugins/bin/$(NAME).so ./plugins/src/$(NAME).go

run: build
	 ./$(BINARY)