.PHONY: build run-server run-server-docker run-client generate certs

build:
	@go build ./...

test:
	@go test ./... -race

run-server:
	@go run ./cmd/server/*.go

run-server-docker:
	@docker build -t winter-server .
	@docker run --rm -p 8080:8080 winter-server

run-client:
	@go run ./cmd/client/*.go

generate:
	@protoc --go_out=plugins=grpc:. ./api/winter/winter.proto

certs:
	@certstrap --depot-path ./certs init --common-name "ca"
	@certstrap --depot-path ./certs request-cert --common-name client
	@certstrap --depot-path ./certs sign client --CA "ca"
	@certstrap --depot-path ./certs request-cert --common-name server
	@certstrap --depot-path ./certs sign server --CA "ca"
	