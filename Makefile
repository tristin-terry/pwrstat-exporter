.PHONY: run build

run:
	go run ./...

build:
	go build ./...

test:
	go test ./...