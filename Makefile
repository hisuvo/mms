.PHONY: run
run:
	air

.PHONY: build
build:
	go build -o bin/mms ./cmd

.PHONY: start
start:
	go run ./cmd