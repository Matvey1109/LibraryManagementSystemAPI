all: build exec

build:
	go build -o bin/main cmd/main/main.go

exec:
	./bin/main

run:
	go run cmd/main/main.go

clear:
	rm -rf bin

help:
	@echo "Available commands:"
	@echo "all    -  Builds and then executes the project"
	@echo "build  -  Builds the project"
	@echo "exec   -  Executes the built binary"
	@echo "run    -  Runs project using 'go run'"
	@echo "clear  -  Cleans the build directory (bin)"
