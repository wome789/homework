package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"math/big"
)

/*
*## 任务 1：区块链读写 任务目标
使用 Sepolia 测试网络实现基础的区块链交互，包括查询区块和发送交易。

	具体任务

1. 环境搭建
  - 安装必要的开发工具，如 Go 语言环境、 go-ethereum 库。
  - 注册 Infura 账户，获取 Sepolia 测试网络的 API Key。

2. 查询区块
  - 编写 Go 代码，使用 ethclient 连接到 Sepolia 测试网络。
  - 实现查询指定区块号的区块信息，包括区块的哈希、时间戳、交易数量等。
  - 输出查询结果到控制台。

3. 发送交易
  - 准备一个 Sepolia 测试网络的以太坊账户，并获取其私钥。
  - 编写 Go 代码，使用 ethclient 连接到 Sepolia 测试网络。
  - 构造一笔简单的以太币转账交易，指定发送方、接收方和转账金额。
  - 对交易进行签名，并将签名后的交易发送到网络。
  - 输出交易的哈希值。
*/
func main1() {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/d59d659eb3ab472d92c4a3a0def1b91d")
	if err != nil {
		fmt.Println("Failed to connect to the Ethereum client: ", err)
		return
	}

	// 获取chainId
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		fmt.Println("Failed to get chain ID: ", err)
		return
	}

	// 从私钥获取转账的privateKey
	privateKey, err := crypto.HexToECDSA("3ee41394c48c7cd0afa0cd6cf9c4aafe79c0da997fbac24f53128f080463070b")
	if err != nil {
		fmt.Println("Failed to get private key: ", err)
		return
	}

	toAddress := common.HexToAddress("0xA8a1e31ac2D41340699E73268E49bec8C28853C3")

	// 获取nonce
	publicKey := privateKey.Public().(*ecdsa.PublicKey)
	fromAddress := crypto.PubkeyToAddress(*publicKey)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		fmt.Println("Failed to get nonce: ", err)
		return
	}

	// 新建一笔转账交易
	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   chainID,
		Nonce:     nonce,
		To:        &toAddress,
		Value:     big.NewInt(100000000000000000), // 0.1 ETH
		GasTipCap: big.NewInt(2 * params.GWei),
		GasFeeCap: big.NewInt(4 * params.GWei),
		Gas:       21000,
	})

	// 对交易进行签名
	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(chainID), privateKey)
	if err != nil {
		fmt.Println("Failed to sign transaction: ", err)
		return
	}

	// 发送交易
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		fmt.Println("Failed to send transaction: ", err)
		return
	}
	// 获取交易hash
	fmt.Println("Transaction hash: ", signedTx.Hash().Hex())

	// 等待交易执行成功
	fmt.Println("Waiting for transaction to be mined...")
	receipt, err := bind.WaitMined(context.Background(), client, signedTx)
	if err != nil {
		fmt.Println("Failed to wait for transaction to be mined: ", err)
		return
	}

	// FIXME 修改地址为上述signTx
	transactionHash := common.HexToHash("0xc34260c63a0b23fd4df5bfa5f85a372cd299b93d56916daeed4b740891bfa084")
	transactionHash = signedTx.Hash()

	/*receipt, err := client.TransactionReceipt(context.Background(), transactionHash)
	if err != nil {
		fmt.Println("Failed to get transaction info: ", err)
		return
	}*/

	// 获取交易的相关信息 查询交易信息
	txInfo, _, err := client.TransactionByHash(context.Background(), transactionHash)
	if err != nil {
		fmt.Println("Failed to get transaction info: ", err)
		return
	}
	// 查询请求头信息
	headerInfo, err := client.HeaderByHash(context.Background(), receipt.BlockHash)
	if err != nil {
		fmt.Println("Failed to get header info: ", err)
		return
	}
	from, err := types.Sender(types.LatestSignerForChainID(chainID), txInfo)
	if err != nil {
		fmt.Println("Failed to get sender address: ", err)
		return
	}

	baseFee := headerInfo.BaseFee
	gasUsed := receipt.GasUsed
	gasLimit := txInfo.Gas()

	maxFee := txInfo.GasFeeCap()
	tipCap := txInfo.GasTipCap()
	// 实际单价
	actualPrice := minBigInt(new(big.Int).Add(baseFee, tipCap), maxFee)
	// 转账费用
	transferFee := new(big.Int).Mul(actualPrice, new(big.Int).SetUint64(gasUsed))
	// 销毁费用
	burnt := new(big.Int).Mul(baseFee, new(big.Int).SetUint64(gasUsed))
	// 节省费用
	savings := new(big.Int).Mul(new(big.Int).Sub(maxFee, actualPrice), new(big.Int).SetUint64(gasUsed))

	fmt.Println("Transaction Hash:", transactionHash)
	fmt.Println("Status:", receipt.Status)
	fmt.Println("Block:", headerInfo.Number)
	fmt.Println("Timestamp:", headerInfo.Time)
	fmt.Println("From:", from)
	fmt.Println("To:", txInfo.To())

	fmt.Println("Transaction Fee:", transferFee)
	fmt.Println("Gas Price:", actualPrice)
	fmt.Println("Gas Limit :", gasLimit)
	fmt.Println("Usage by Txn:", gasUsed)
	fmt.Println("BaseFee:", baseFee)
	fmt.Println("Max:", maxFee)
	fmt.Println("Max Priority:", tipCap)
	fmt.Println("Burnt:", burnt)
	fmt.Println("Txn Savings:", savings)
	fmt.Println("Txn Type:", txInfo.Type())
	fmt.Println("Nonce:", txInfo.Nonce())
	fmt.Println("Position In Block:", receipt.TransactionIndex)

	// 获取区块信息
	block, err := client.BlockByHash(context.Background(), receipt.BlockHash)
	if err != nil {
		fmt.Println("Failed to get block info: ", err)
		return
	}

	// 实现查询指定区块号的区块信息，包括区块的哈希、时间戳、交易数量等。
	fmt.Println("block Hash:", block.Hash())
	fmt.Println("block time:", block.Time())
	fmt.Println("block trans count:", block.Transactions().Len())

}

func minBigInt(a, b *big.Int) *big.Int {
	if a.Cmp(b) < 0 {
		return a
	}
	return b
}
