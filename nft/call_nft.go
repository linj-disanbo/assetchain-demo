


/*

import (
    "context"
    "fmt"
    "math/big"

    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/crypto"
)

func main() {
    client, err := ethclient.Dial("https://ropsten.infura.io")
    if err != nil {
        panic(err)
    }

    contractAddress := common.HexToAddress("0x123...")
    instance, err := NewContract(contractAddress, client)
    if err != nil {
        panic(err)
    }

    // ...
}


func main() {
	// 连接以太坊节点
	client, err := ethclient.Dial(env.RPC)
	if err != nil {
		panic(err)
	}

	// 合约地址
	contractAddr := common.HexToAddress(env.NftContract)

	// 合约ABI定义
	abiDef, err := abi.JSON([]byte(nft.AsNFTABI))
	if err != nil {
		panic(err)
	}

	// 调用mintNFT方法
	recipient := common.HexToAddress(env.TestAddress2)
	asCode := "c1i4zs4MgFi"
	tx, err := callContract(client, contractAddr, abiDef, "mintNFT", recipient, asCode)
	if err != nil {
		panic(err)
	}
	fmt.Println(tx.Hash().Hex())
}

func callContract(client *ethclient.Client, contract common.Address, abi abi.ABI, method string, args ...interface{}) (*ethereum.Transaction, error) {
	// 获取nonce值
	nonce, err := client.PendingNonceAt(context.Background(), common.HexToAddress("0x1234567890123456789012345678901234567890"))
	if err != nil {
		return nil, err
	}

	// 构造调用合约的交易
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	txInput, err := abi.Pack(method, args...)
	if err != nil {
		return nil, err
	}

	tx := ethereum.Transaction{
		From:     common.HexToAddress("0x1234567890123456789012345678901234567890"),
		To:       &contract,
		Nonce:    nonce,
		GasPrice: gasPrice,
		GasLimit: 300000, // 可以通过estimateGas方法进行估算
		Value:    nil,
		Data:     txInput,
	}

	// 对交易进行签名
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(env.ID), env.TestPrivkey)
	if err != nil {
		return nil, err
	}

	// 发送交易
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return nil, err
	}

	return signedTx, nil
}
*/