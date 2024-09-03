PORT_CONFIG := 65432:5432

all: run

build:
	@echo "Building binary..."
	@go build -o bin/main .

run:
	@echo "Starting webserver..."
	@air

database:
	@echo "Creating database..."
	@docker compose run -p $(PORT_CONFIG) -d db 

test:
	@echo "Running tests..."
	@go test -v ./test

.PHONY: clean test
clean:
	@rm -rf bin
