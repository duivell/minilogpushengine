.PHONY: build run debug test vet clean

build:
	go build -o bin/minilogpush ./cmd/minilogpush

run:
	go run ./cmd/minilogpush

# Headless so a DAP client (VS Code, etc.) can attach on :2345.
# For a plain terminal debugger instead, run: dlv debug ./cmd/minilogpush
debug:
	dlv debug ./cmd/minilogpush --headless --listen=:2345 --api-version=2 --accept-multiclient

test:
	go test ./...

vet:
	go vet ./...

clean:
	rm -rf bin
