.PHONY: start build

EXPORT_FOLDER = "export"
SERVER_NAME = "little-bee-server"

all: start

build:
	@go build -ldflags "-w -s" -o ./$(EXPORT_FOLDER)/$(SERVER_NAME) ./cmd/server

start:
	go run cmd/server/main.go -c ./configs/config.toml

clean:
	rm -rf $(EXPORT_FOLDER)