.PHONY: hotreload lint test whatsapp worker job

hotreload:
	air --build.cmd "go build -o bin/whatsapp cmd/whatsapp/main.go" --build.bin "./bin/whatsapp"

lint:
	golangci-lint run ./...

test:
	go test ./internal/tests/...

whatsapp:
	cp .env.whatsapp .env && go run cmd/whatsapp/main.go

job:
	cp .env.job .env && go run cmd/job/main.go
