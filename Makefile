GOPATH=$(shell go env GOPATH)
LDFLAGS="-w"

.PHONY: default
default: lint test

.PHONY: lint
lint:
	go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.37.0
	$(GOPATH)/bin/golangci-lint run -e gosec ./...
	go fmt ./...
	go mod tidy

# added -race in future (badger fatal error: checkptr: pointer arithmetic result points to invalid allocation)
# https://github.com/golang/go/issues/40917
.PHONY: test
test:
	go test *.go

.PHONY: build
build:
	go build -ldflags=$(LDFLAGS) -o vpn *.go
