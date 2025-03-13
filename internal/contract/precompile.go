package contract

import (
	"bufio"
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math/big"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func CompileContract(filename string) error {
	cmd := exec.Command("solc", "--abi", "--bin", "-o", "abi", filename)
	_, err := cmd.Output()
	if err != nil {
		return err
	}

	return nil
}

func GetContractBytes(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var bytecode string
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "Binary:") {
			bytecode = line
		}
	}

	if bytecode == "" {
		return nil, fmt.Errorf("Bytecode not found")
	}

	bytecodeBytes, err := hex.DecodeString(bytecode)
	if err != nil {
		return nil, err
	}

	return bytecodeBytes, nil
}

func DeployContract(client *ethclient.Client, privateKey *ecdsa.PrivateKey, bytecode []byte) (common.Hash, error) {
	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return common.Hash{}, fmt.Errorf("Failed to get nonce: %v", err)
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return common.Hash{}, fmt.Errorf("Failed to get chain ID: %v", err)
	}

	tx := types.NewContractCreation(nonce, big.NewInt(0), uint64(3000000), big.NewInt(1e9), bytecode)

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return common.Hash{}, fmt.Errorf("Failed to sign transaction: %v", err)
	}

	if err = client.SendTransaction(context.Background(), signedTx); err != nil {
		return common.Hash{}, fmt.Errorf("Failed to send transaction: %v", err)
	}

	if _, err = WaitForReceipt(client, signedTx.Hash()); err != nil {
		return common.Hash{}, fmt.Errorf("Failed to get receipt: %v", err)
	}

	return signedTx.Hash(), nil
}

func WaitForReceipt(client *ethclient.Client, txHash common.Hash) (*types.Receipt, error) {
	for {
		receipt, err := client.TransactionReceipt(context.Background(), txHash)
		if err == nil {
			return receipt, nil
		}
		if err != ethereum.NotFound {
			return nil, err
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func GenerateAbi(abi string, bin string) error {
	cmd := exec.Command("abigen", "--abi", abi, "--bin", bin, "--pkg", "precompile", "--type", "Precompile", "--out", "precompile/precompile.go")
	return cmd.Run()
}
