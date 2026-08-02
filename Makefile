.PHONY: fmt tidy test lint check

fmt:
	gofmt -w .

tidy:
	go mod tidy

# All tests, no skips — includes the hq structural lint (internal/hqlint).
test:
	go test ./...

lint:
	golangci-lint run

# The one gate to run before every commit: everything green.
check: fmt tidy test lint
