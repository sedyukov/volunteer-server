.PHONY: start

build:
	(go build ./cmd/gateway)

start:
	(cd ./cmd/gateway && go run .)

lint:
	golangci-lint run
