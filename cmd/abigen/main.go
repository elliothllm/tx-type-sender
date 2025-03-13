package main

import (
	"fmt"
	"log"

	"github.com/elliothllm/tx-type-sender/internal/contract"
)

func main() {
	fmt.Println("Compiling the contract...")
	err := contract.CompileContract("contracts/precompile.sol")
	if err != nil {
		log.Fatalf("Failed to compile contract: %v", err)
	}

	fmt.Println("Generating ABI bindings...")
	err = contract.GenerateAbi("abi/PrecompileTester.abi", "abi/PrecompileTester.bin")
	if err != nil {
		log.Fatalf("Failed to generate ABI bindings: %v", err)
	}

	fmt.Println("ABI bindings generated successfully.")
}
