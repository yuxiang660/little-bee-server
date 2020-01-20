.PHONY: start build

SERVER_NAME = "little-bee-server"

all: start

build:
	@go build -ldflags "-w -s" -o $(SERVER_NAME) ./cmd/server

start:
	go run cmd/server/main.go -c ./configs/config.toml

clean:
	rm -rf $(SERVER_NAME)