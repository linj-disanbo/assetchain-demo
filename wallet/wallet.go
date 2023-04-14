package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/linj-disanbo/assetchain-demo/env"
)

/*

上面代码中，我们使用了 Ethereum Go客户端库（ethclient）来连接以太坊节点，然后使用 client.BalanceAt 函数查询了指定地址的余额。最后将以太转换为最小单位Wei并输出余额。

Ethereum Go客户端库和 Web3.js 都提供了一个完整的API，能够访问以太坊的各个部分。除了余额查询之外，您还可以使用这些库来发送交易、调用合约、查询区块链历史记录等。
*/

func main() {
	// 以太坊节点的URL
	url := env.RPC
	// 创建以太坊客户端
	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}
	// 要查询的地址
	address := common.HexToAddress(env.TestAddress2)
	// 查询余额
	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatal(err)
	}
	// 将以太转换为最小单位wei
	ethBalance := new(big.Float).Quo(new(big.Float).SetInt(balance), big.NewFloat(params.Ether))
	// 输出余额
	fmt.Printf("Balance: %s ETH\n", ethBalance.Text('f', 4))
}
