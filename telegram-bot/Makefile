SHELL=/bin/bash
BINARY_NAME=main.bin

prerequisites:
	go mod tidy

clean:
	go clean
	rm -f "$(BINARY_NAME)" __debug_bin

build: prerequisites
	go build -o $(BINARY_NAME) cmd/main.go

run: build
	bash -c 'BOT_TOKEN=$(TOKEN) BOT_CHAT_ID=$(CHAT_ID) ./$(BINARY_NAME)'
