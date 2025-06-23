BINARY_NAME=hassio-api
MAIN=cmd/api/main.go

.PHONY: run generate build dev

run: generate
	go run $(MAIN)

build: generate
	go build -o bin/$(BINARY_NAME) $(MAIN)

generate:
	go generate .

dev:
	go run -mod=mod github.com/air-verse/air
