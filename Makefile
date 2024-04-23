all: build exec

build:
	go build -o bin/main cmd/main/main.go

exec:
	./bin/main

run:
	go run cmd/main/main.go

test:
	go test -v ./internal/tests

clear:
	rm -rf bin

dock:
	docker compose up -d

help:
	@echo "Available commands:"
	@echo "all    -  Builds and then executes the project"
	@echo "build  -  Builds the project"
	@echo "exec   -  Executes the built binary"
	@echo "run    -  Runs project using 'go run'"
	@echo "test   -  Test project"
	@echo "clear  -  Cleans the build directory (bin)"
	@echo "dock   -  Dockerize app and run"
