PID_FILE := .minilogpush.pid
LOG_FILE := minilogpush.log

.PHONY: build run start stop restart debug test vet clean

build:
	go build -o bin/minilogpush ./cmd/minilogpush

run:
	go run ./cmd/minilogpush

# Runs the built binary in the background; logs go to $(LOG_FILE).
start: build
	@if [ -f $(PID_FILE) ] && kill -0 $$(cat $(PID_FILE)) 2>/dev/null; then \
		echo "already running (pid $$(cat $(PID_FILE)))"; \
	else \
		(./bin/minilogpush > $(LOG_FILE) 2>&1 & echo $$! > $(PID_FILE)); \
		echo "started (pid $$(cat $(PID_FILE)))"; \
	fi

stop:
	@if [ -f $(PID_FILE) ] && kill -0 $$(cat $(PID_FILE)) 2>/dev/null; then \
		kill $$(cat $(PID_FILE)) && rm -f $(PID_FILE) && echo "stopped"; \
	else \
		echo "not running"; rm -f $(PID_FILE); \
	fi

# Rebuilds (via start's dependency) and swaps in the new binary.
restart: stop start

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
