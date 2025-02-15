.PHONY: test
test:
	go test -v -race ./...

.PHONY: lint
lint:
	go tool github.com/golangci/golangci-lint/cmd/golangci-lint run
