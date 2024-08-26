all: run

build:
	@echo "Building binary..."
	@go build -o bin/main .

run:
	@echo "Starting webserver..."
	@air

test:
	@echo "Running tests..."
	@go test -v ./test

.PHONY: clean test
clean:
	@rm -rf bin
