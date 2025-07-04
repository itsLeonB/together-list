.PHONY: hotreload lint test whatsapp worker

hotreload:
	air --build.cmd "go build -o bin/whatsapp cmd/whatsapp/main.go" --build.bin "./bin/whatsapp"

lint:
	golangci-lint run ./...

test:
	go test ./internal/tests/...

whatsapp:
	go run cmd/whatsapp/main.go

worker:
	go run cmd/worker/main.go
