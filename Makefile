.PHONY: start

build:
	(go build ./cmd/server)

start:
	(cd ./cmd/server && go run .)

lint:
	golangci-lint run
