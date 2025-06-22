BINARY_NAME=hassio-api
MAIN=cmd/api/main.go

.PHONY: run generate entgql build clean dbup

run: generate
	go run $(MAIN)

build: generate
	go build -o bin/$(BINARY_NAME) $(MAIN)

generate:
	go generate .

# очистка
clean:
	rm -rf ./ent/generated ./graph/generated ./graph/models.go ./bin

dbup:
	docker-compose up -d db

dev:
	go run -mod=mod github.com/air-verse/air
