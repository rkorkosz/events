build:
	CGO_ENABLED=0 go build -ldflags="-s -w" -o events ./cmd/*.go

vet:
	go vet ./cmd/... ./pkg/...

golangci:
	golangci-lint run

test: vet golangci
	go test -race -v ./...
