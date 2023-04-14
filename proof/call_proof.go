package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/linj-disanbo/assetchain-demo/env"
	"github.com/linj-disanbo/assetchain-demo/proof/proof"
)

func main() {
	client, err := ethclient.Dial(env.RPC)
	if err != nil {
		panic(err)
	}

	contractAddress := common.HexToAddress(env.ProofContract)
	instance, err := proof.NewMyProof(contractAddress, client)
	if err != nil {
		panic(err)
	}

	privateKey, err := crypto.HexToECDSA(env.TestPrivkey)
	if err != nil {
		panic(err)
	}

	from := common.HexToAddress(env.TestAddress)
	nonce, err := client.PendingNonceAt(context.Background(), from)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(int64(env.ID)))
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasPrice = gasPrice
	auth.GasLimit = uint64(300000)
	auth.Value = big.NewInt(0)

	hash := "hash(as-code-no1113)"

	tx, err := instance.SaveProof(auth, hash)
	if err != nil {
		panic(err)
	}

	fmt.Println("Transaction hash:", tx.Hash().Hex())

	/*
		// Wait for the transaction to be mined
		ctx := context.Background()
		tx2, err := bind.WaitMined(ctx, client, tx)
		if err != nil {
			panic(err)
		}

		fmt.Println("Transaction hash:", tx2.TxHash.Hex())

		for _, event := range tx2.Logs {
			fmt.Println(event)
		}
	*/
}
