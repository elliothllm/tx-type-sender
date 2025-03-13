package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/elliothllm/tx-type-sender/cmd"
	"github.com/elliothllm/tx-type-sender/internal/contract"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/elliothllm/tx-type-sender/precompile"
)

func main() {
	ctx := context.Background()

	file, err := os.Open("sender-cfg.json")
	if err != nil {
		fmt.Println("error opening file:", err)
		return
	}
	defer file.Close()

	var cfg *cmd.Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	if err = testPrecompiles(ctx, cfg); err != nil {
		fmt.Println("Error testing precompiles:", err)
		return
	}
}

func testPrecompiles(ctx context.Context, cfg *cmd.Config) error {
	client, err := ethclient.Dial(cfg.RpcURL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	privateKey, err := crypto.HexToECDSA(cfg.FromPrivateKey)
	if err != nil {
		log.Fatalf("Failed to create private key: %v", err)
	}

	chainID, err := client.ChainID(ctx)
	if err != nil {
		log.Fatalf("Failed to get the chain ID: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatalf("Failed to create transactor: %v", err)
	}

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		log.Fatalf("Failed to suggest gas price: %v", err)
	}
	auth.GasPrice = gasPrice
	auth.GasLimit = uint64(3000000)

	address, tx, contractInstance, err := precompile.DeployPrecompile(auth, client)
	if err != nil {
		log.Fatalf("Failed to deploy contract: %v", err)
	}

	fmt.Printf("Contract deployment initiated:\n  Address: %s\n  Tx Hash: %s\n", address.Hex(), tx.Hash().Hex())

	receipt, err := contract.WaitForReceipt(client, tx.Hash())
	if err != nil {
		log.Fatalf("Error waiting for receipt: %v", err)
	}

	fmt.Printf("Contract deployed in block: %d\n", receipt.BlockNumber.Uint64())

	contractInstance, err = precompile.NewPrecompile(address, client)
	if err != nil {
		log.Fatalf("Failed to create contract binding: %v", err)
	}

	callOpts := &bind.CallOpts{Context: ctx}

	// ecrecover (address 0x01)
	ecrecoverRes, err := contractInstance.TestEcrecover(callOpts, []byte("hello world"))
	if err != nil {
		log.Fatalf("TestEcrecover failed: %v", err)
	} else {
		fmt.Printf("TestEcrecover succeeded, TestEcrecover('hello world') = %x\n", ecrecoverRes)
	}

	// sha256 (address 0x02)
	sha256Res, err := contractInstance.TestSha256(callOpts, []byte("hello world"))
	if err != nil {
		log.Fatalf("TestSha256 failed: %v", err)
	} else {
		fmt.Printf("TestSha256 succeeded, TestSha256('hello world') = %x\n", sha256Res)
	}

	// ripemd160 (address 0x03)
	ripemd160Res, err := contractInstance.TestRipemd160(callOpts, []byte("hello world"))
	if err != nil {
		log.Fatalf("TestRipemd160 failed: %v", err)
	} else {
		fmt.Printf("TestRipemd160 succeeded, TestRipemd160('hello world') = %x\n", ripemd160Res)
	}

	// identity (address 0x04)
	identityRes, err := contractInstance.TestIdentity(callOpts, []byte("hello world"))
	if err != nil {
		log.Fatalf("TestIdentity failed: %v", err)
	} else {
		fmt.Printf("TestIdentity succeeded, TestIdentity('hello world') = %x\n", identityRes)
	}

	// modexp (address 0x05)
	modexpRes, err := contractInstance.TestModexp(callOpts, []byte("hello world"))
	if err != nil {
		log.Fatalf("TestModexp failed: %v", err)
	} else {
		fmt.Printf("TestModexp succeeded, TestModexp('hello world') = %x\n", modexpRes)
	}

	// bn256Add (address 0x06)
	bn256AddRes, err := contractInstance.TestBn256Add(callOpts, []byte("hello world"))
	if err != nil {
		log.Fatalf("TestBn256Add failed: %v", err)
	} else {
		fmt.Printf("TestBn256Add succeeded, TestBn256Add('hello world') = %x\n", bn256AddRes)
	}

	// bn256ScalarMul (address 0x07)
	bn256ScalarMulRes, err := contractInstance.TestBn256ScalarMul(callOpts, []byte("hello world"))
	if err != nil {
		log.Fatalf("TestBn256ScalarMul failed: %v", err)
	} else {
		fmt.Printf("TestBn256ScalarMul succeeded, TestBn256ScalarMul('hello world') = %x\n", bn256ScalarMulRes)
	}

	// bn256Pairing (address 0x08)
	_, err = contractInstance.TestBn256Pairing(callOpts, []byte("hello world"))
	if err != nil {
		log.Fatalf("TestBn256Pairing failed: %v", err)
	} else {
		fmt.Printf("TestBn256Pairing succeeded")
	}

	// blake2 (address 0x09)
	blake2Res, err := contractInstance.TestBlake2(callOpts, []byte("hello world"))
	if err != nil {
		log.Fatalf("TestBlake2 failed: %v", err)
	} else {
		fmt.Printf("TestBlake2 succeeded, TestBlake2('hello world') = %x\n", blake2Res)
	}

	// pointEvaluation (address 0x0A)
	pointEvaluationRes, err := contractInstance.TestPointEvaluation(callOpts, []byte("hello world"))
	if err != nil {
		log.Fatalf("TestPointEvaluation failed: %v", err)
	} else {
		fmt.Printf("TestPointEvaluation succeeded, TestPointEvaluation('hello world') = %x\n", pointEvaluationRes)
	}

	// BLS12_G1ADD (address 0x0B)
	_, err = contractInstance.TestBLS12G1Add(callOpts, []byte("hello world"))
	if err != nil {
		log.Fatalf("TestBLS12G1Add failed: %v", err)
	} else {
		fmt.Printf("TestBLS12G1Add succeeded")
	}

	// BLS12_G1MSM (address 0x0C)
	_, err = contractInstance.TestBLS12G1MSM(callOpts, []byte("hello world"))
	if err != nil {
		log.Fatalf("TestBLS12G1MSM failed: %v", err)
	} else {
		fmt.Printf("TestBLS12G1MSM succeeded")
	}

	// BLS12_G2ADD (address 0x0D)
	_, err = contractInstance.TestBLS12G2Add(callOpts, []byte("hello world"))
	if err != nil {
		log.Fatalf("TestBLS12G2Add failed: %v", err)
	} else {
		fmt.Printf("TestBLS12G2Add succeeded")
	}

	// BLS12_G2MSM (address 0x0E)
	_, err = contractInstance.TestBLS12G2MSM(callOpts, []byte("hello world"))
	if err != nil {
		log.Fatalf("TestBLS12G2MSM failed: %v", err)
	} else {
		fmt.Printf("TestBLS12G2MSM succeeded")
	}

	// BLS12_PAIRING_CHECK (address 0x0F)
	_, err = contractInstance.TestBLS12PairingCheck(callOpts, []byte("hello world"))
	if err != nil {
		log.Fatalf("TestBLS12PairingCheck failed: %v", err)
	} else {
		fmt.Printf("TestBLS12PairingCheck succeeded")
	}

	// BLS12_MAP_FP_TO_G1 (address 0x10)
	_, err = contractInstance.TestBLS12MapFPtoG1(callOpts, []byte("hello world"))
	if err != nil {
		log.Fatalf("TestBLS12MapFPtoG1 failed: %v", err)
	} else {
		fmt.Printf("TestBLS12MapFPtoG1 succeeded")
	}

	// BLSS12_MAP_FP2_TO_G2 (address 0x11)
	_, err = contractInstance.TestBLSS12MapFP2toG2(callOpts, []byte("hello world"))
	if err != nil {
		log.Fatalf("TestBLSS12MapFP2toG2 failed: %v", err)
	} else {
		fmt.Printf("TestBLSS12MapFP2toG2 succeeded")
	}

	return nil
}
