run:
	@go run cmd/cli/main.go

build:
	@cd cmd/cli && go build -o ../../bin/cli

build-clean:
	@rm -rf ./bin/*

vet:
	@go vet ./cmd/cli