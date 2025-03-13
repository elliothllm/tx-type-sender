.PHONY: send precompile abigen

send:
	@go run cmd/sender/main.go

precompile:
	@go run cmd/precompile/main.go

abigen:
	@rm -rf abi/*
	@rm -rf precompile/precompile.go
	@go run cmd/abigen/main.go

solc-install:
	@go install github.com/ethereum/go-ethereum/cmd/solc@latest

abigen-install:
	@go install github.com/ethereum/go-ethereum/cmd/abigen@latest
