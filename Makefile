.PHONY: start

build:
	(CGO_ENABLED=0 go build ./cmd/server)

start:
	(cd ./cmd/server && go run .)

lint:
	golangci-lint run

deploy:
	(CGO_ENABLED=0 go build ./cmd/server && scp server root@example:~/server)