PHONY: build
build:
	go build -o auto-messenger backend/cmd/server/main.go

PHONY: run
run:
	./auto-messenger

PHONY: hot
hot:
	go run backend/cmd/server/main.go
