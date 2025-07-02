.PHONY: hotreload lint

hotreload:
	air --build.cmd "go build -o bin/whatsapp cmd/whatsapp/main.go" --build.bin "./bin/whatsapp"

lint:
	golangci-lint run ./...

test:
	go test ./internal/tests/...
