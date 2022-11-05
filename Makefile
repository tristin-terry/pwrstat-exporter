.PHONY: run build local-install

run:
	go run ./... --listen-address 10101

build:
	go build ./cmd/pwrstat-exporter/

test:
	go test ./... -cover

test-ci:
	go test ./... -race -coverprofile=coverage.out -covermode=atomic

local-install: build 
	mv pwrstat-exporter /usr/local/bin/ && systemctl restart pwrstat-exporter.service