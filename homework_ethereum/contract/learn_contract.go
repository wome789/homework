package main

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"time"
	"transfer/contract/my_count"
)

/*
*

	任务 2：合约代码生成 任务目标:
		使用 abigen 工具自动生成 Go 绑定代码，用于与 Sepolia 测试网络上的智能合约进行交互。

	具体任务:
	编写智能合约
	使用 Solidity 编写一个简单的智能合约，例如一个计数器合约。
	编译智能合约，生成 ABI 和字节码文件。
	使用 abigen 生成 Go 绑定代码
	安装 abigen 工具。
	使用 abigen 工具根据 ABI 和字节码文件生成 Go 绑定代码。
	使用生成的 Go 绑定代码与合约交互
	编写 Go 代码，使用生成的 Go 绑定代码连接到 Sepolia 测试网络上的智能合约。
	调用合约的方法，例如增加计数器的值。
	输出调用结果。
*/
func main() {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/d59d659eb3ab472d92c4a3a0def1b91d")
	if err != nil {
		fmt.Println("Failed to connect to the Ethereum client:", err)
		return
	}

	// 获取部署者私钥对象
	privateKey, _ := crypto.HexToECDSA("a2082c324f3a54494cbd1854436938549d98160fae2512dffa6cca92336d6e61")
	// 获取部署者地址
	publicKey := privateKey.Public().(*ecdsa.PublicKey)
	deployerAddress := crypto.PubkeyToAddress(*publicKey)

	// 生成auth对象
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(11155111))
	if err != nil {
		fmt.Println("Failed to create authorized transactor:", err)
		return
	}

	// 部署合约
	// address:0x66b7921f558cA25FfcB727093c25549c6a2E95Ce
	// tx:0x353aa5e066a9d1062cfe33ebfd88f75cd43cac523b8e3614a4b45157231a8b53
	myCountAddress, transaction, myCountInterface, err := my_count.DeployMyCount(auth, client)
	if err != nil {
		fmt.Println("Failed to deploy contract:", err)
		return
	}

	fmt.Println("Contract deployed at address:", myCountAddress.Hex())
	fmt.Println("Transaction hash:", transaction.Hash().Hex())
	fmt.Println("MyCount interface:", myCountInterface)

	receipt, err := waitTxFinish(client, transaction.Hash())
	if err != nil {
		fmt.Println("Failed to get transaction receipt:", err)
		return
	}
	fmt.Println("Transaction confirmed in block:", receipt.BlockNumber)

	count, err := myCountInterface.GetCount(&bind.CallOpts{}, deployerAddress)
	if err != nil {
		fmt.Println("Failed to get count:", err)
		return
	}
	fmt.Println("Count:", count)

	// 发送设置 0x2f260e3b4f1a3d2c4e8204d13caa2b4d447fd43c70a1531d206eaccfb110a6e4
	transaction, err = myCountInterface.Increment(auth, deployerAddress)
	if err != nil {
		fmt.Println("Failed to increment count:", err)
		return
	}
	fmt.Println("Increment transaction hash:", transaction.Hash().Hex())
	receipt, err = waitTxFinish(client, transaction.Hash())
	if err != nil {
		fmt.Println("Failed to get transaction receipt:", err)
		return
	}
	// 重新获取
	count, err = myCountInterface.GetCount(&bind.CallOpts{}, deployerAddress)
	if err != nil {
		fmt.Println("Failed to get count:", err)
		return
	}
	fmt.Println("Count2:", count)
}

func waitTxFinish(client *ethclient.Client, tx common.Hash) (*types.Receipt, error) {
	i := 0
	for {
		receipt, err := client.TransactionReceipt(context.Background(), tx)
		if err == nil {
			return receipt, nil
		}
		if !errors.Is(err, ethereum.NotFound) {
			return nil, err
		}
		i++
		fmt.Println("Transaction not found, waiting...", i)
		time.Sleep(1 * time.Second)
	}
}
