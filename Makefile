MIGRATIONS_PATH="configs/migrations"
MIGRATION_NAME="MIGRATION_NAME"


all: build test run

install:
	go get -u -v -f all
	go mod tidy

test:
	go test -v ./... | grep -i fail

run:
	go run ./cmd/go_pg_poc.go

build:
	go build ./cmd/go_pg_poc.go
