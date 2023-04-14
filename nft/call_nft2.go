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
	"github.com/linj-disanbo/assetchain-demo/nft/nft"
)

func main() {
	client, err := ethclient.Dial(env.RPC)
	if err != nil {
		panic(err)
	}

	contractAddress := common.HexToAddress(env.NftContract)
	instance, err := nft.NewAsNFT(contractAddress, client)
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

	recipient := common.HexToAddress(env.TestAddress2)
	asCode := "as-code-no1111"

	tx, err := instance.MintNFT(auth, recipient, asCode)
	if err != nil {
		panic(err)
	}

	fmt.Println("Transaction hash:", tx.Hash().Hex())
}
