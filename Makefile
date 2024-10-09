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

.PHONY: clean test seed
test:
	@echo "Running tests..."
	@go test -v ./test

seed:
	@echo "Creating seed data..."
	@./seed/characters/seed-characters.sh
	@./seed/songs/get-songs.sh

clean:
	@rm -rf bin
