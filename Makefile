.PHONY: run build local-install

run:
	go run ./...

build:
	go build ./cmd/pwrstat-exporter/

test:
	go test ./...

local-install: build 
	mv pwrstat-exporter /usr/local/bin/ && systemctl restart pwrstat-exporter.service