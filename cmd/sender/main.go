package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/elliothllm/tx-type-sender/cmd"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/holiman/uint256"
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

	if cfg.Type < 1 || cfg.Type > 4 {
		log.Fatalf("Invalid transaction type: %d Must be 1/2/3/4", cfg.Type)
	}

	for i := 0; i < cfg.TxAmount; i++ {
		sendTxs(ctx, cfg)
		time.Sleep(time.Duration(cfg.TxSendIntervalSeconds) * time.Second)
	}
}

func sendTxs(ctx context.Context, cfg *cmd.Config) {

	client, err := ethclient.Dial(cfg.RpcURL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	privateKey, err := crypto.HexToECDSA(cfg.FromPrivateKey)
	if err != nil {
		log.Fatalf("Failed to create private key: %v", err)
	}

	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)

	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		log.Fatalf("Failed to get the nonce: %v", err)
	}

	toAddress := common.HexToAddress(cfg.ToAddress)

	chainID, err := client.ChainID(ctx)
	if err != nil {
		log.Fatalf("Failed to get the chain ID: %v", err)
	}

	var tx *types.Transaction

	switch cfg.Type {
	case 1:
		fmt.Println("Using type 1 Legacy transaction")

		tx = types.NewTx(&types.LegacyTx{
			Nonce:    nonce,
			GasPrice: big.NewInt(int64(cfg.GasPrice)),
			Gas:      cfg.GasLimit,
			To:       &toAddress,
			Value:    big.NewInt(int64(cfg.SendValue)),
			Data:     []byte{},
		})
	case 2:
		fmt.Println("Using Type 2 Access List transaction (EIP-2930)")

		tx = types.NewTx(&types.AccessListTx{
			ChainID:    chainID,
			Nonce:      nonce,
			GasPrice:   big.NewInt(int64(cfg.GasPrice)),
			Gas:        cfg.GasLimit,
			To:         &toAddress,
			Value:      big.NewInt(int64(cfg.SendValue)),
			Data:       []byte{},
			AccessList: types.AccessList{},
		})
	case 3:
		fmt.Println("Using Type 3 Dynamic Fee transaction (EIP-1559)")

		tx = types.NewTx(&types.DynamicFeeTx{
			ChainID:    chainID,
			Nonce:      nonce,
			GasTipCap:  big.NewInt(int64(cfg.GasPrice)),
			GasFeeCap:  new(big.Int).Add(new(big.Int).SetUint64(cfg.GasPrice), new(big.Int).SetUint64(cfg.GasPrice)),
			Gas:        cfg.GasLimit,
			To:         &toAddress,
			Value:      big.NewInt(int64(cfg.SendValue)),
			Data:       []byte{},
			AccessList: types.AccessList{},
		})
	case 4:
		fmt.Println("Using Type 4 Set Code transaction (EIP-7702)")

		tx = types.NewTx(&types.SetCodeTx{
			ChainID:    uint256.NewInt(chainID.Uint64()),
			Nonce:      nonce,
			GasTipCap:  uint256.NewInt(cfg.GasPrice),
			GasFeeCap:  uint256.NewInt(cfg.GasPrice * 2),
			Gas:        cfg.GasLimit,
			To:         toAddress,
			Value:      uint256.NewInt(cfg.SendValue),
			Data:       []byte{},
			AccessList: types.AccessList{},
			AuthList:   []types.SetCodeAuthorization{},
		})
	}

	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(chainID), privateKey)
	if err != nil {
		log.Fatalf("Failed to sign transaction: %v", err)
	}

	balanceFrom, err := client.BalanceAt(ctx, fromAddress, nil)
	if err != nil {
		log.Fatalf("Failed to get the balance: %v", err)
	}

	balanceTo, err := client.BalanceAt(ctx, toAddress, nil)
	if err != nil {
		log.Fatalf("Failed to get the balance: %v", err)
	}

	if err = client.SendTransaction(ctx, signedTx); err != nil {
		log.Fatalf("Failed to send the transaction: %v", err)
	}

	var receipt *types.Receipt

	if cfg.WaitForReceipt {
		receipt, err = bind.WaitMined(ctx, client, signedTx)
		if err != nil {
			log.Fatalf("Failed to get transaction receipt: %v", err)
		}
	}

	fmt.Printf("<-----------------NEW TRANSACTION------------------>\n")
	fmt.Printf("Balance of from address: %s\n", balanceFrom)
	fmt.Printf("Balance of to address: %s\n", balanceTo)
	fmt.Printf("Send value: %d\n", cfg.SendValue)
	fmt.Printf("Nonce: %d\n", nonce)
	fmt.Printf("GasPrice: %d\n", cfg.GasPrice)
	fmt.Printf("GasLimit: %d\n", cfg.GasLimit)
	fmt.Printf("Chain ID: %d\n", chainID)
	fmt.Printf("Transaction hash (pre-send): %s\n", signedTx.Hash().Hex())
	fmt.Printf("Transaction sent! TX Hash: %s\n", signedTx.Hash().Hex())
	if cfg.WaitForReceipt {
		fmt.Printf("Gas Used: %d\n", receipt.GasUsed)
		fmt.Printf("Transaction Hash: %s\n", receipt.TxHash.Hex())
		fmt.Printf("Block Number: %s\n", receipt.BlockNumber.String())
		fmt.Printf("Total TX Cost (ETH): %.18f\n", new(big.Float).Quo(new(big.Float).SetInt(new(big.Int).Mul(big.NewInt(int64(cfg.GasPrice)), big.NewInt(int64(receipt.GasUsed)))), big.NewFloat(1e18)))
		fmt.Printf("Status: %d\n", receipt.Status)
	}
}
