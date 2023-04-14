package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/linj-disanbo/assetchain-demo/env"
)

func main() {
	client, err := ethclient.Dial(env.RPC)
	if err != nil {
		log.Fatal(err)
	}

	// sender and receiver addresses
	senderAddress := common.HexToAddress(env.TestAddress)
	receiverAddress := common.HexToAddress(env.TestAddress2)

	// private key of the sender address
	privateKey, err := crypto.HexToECDSA(env.TestPrivkey)
	if err != nil {
		log.Fatal(err)
	}

	// get nonce of the sender address
	nonce, err := client.PendingNonceAt(context.Background(), senderAddress)
	if err != nil {
		log.Fatal(err)
	}

	// gas price and limit
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	gasLimit := uint64(22000)

	// transfer amount
	value := big.NewInt(1000000000000000000) // 1 ETH

	// create transaction
	tx := types.NewTransaction(nonce, receiverAddress, value, gasLimit, gasPrice, nil)

	// sign transaction
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// send transaction
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	// print transaction hash
	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}
