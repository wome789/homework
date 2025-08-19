# NFT拍卖系统 (NFT Auction System)

基于以太坊的NFT拍卖系统，支持多种货币支付和预言机价格转换。

## 📋 项目概述

本项目实现了一个完整的NFT拍卖系统，包含：
- **NFT合约**: 基于ERC721标准的NFT代币
- **拍卖合约**: 支持多种货币支付的拍卖逻辑
- **工厂合约**: 可升级的拍卖合约工厂
- **预言机集成**: Chainlink预言机价格转换
- **代理升级**: UUPS代理模式支持合约升级

## 🏗️ 项目结构

```
homework_hardhat/
├── contracts/
│   ├── nft/ArrowKingNft.sol           # ERC721 NFT合约
│   ├── auction/AuctionContract.sol    # 拍卖合约
│   └── factory/                       # 工厂合约
│       ├── AuctionFactoryContract.sol     # V1版本
│       └── AuctionFactoryContractV2.sol   # V2版本（升级版）
├── deploy/                            # 部署脚本
├── test/AuctionTest.js                # 测试文件
└── hardhat.config.js                  # 配置文件
```

## 🚀 核心功能

### 1. NFT合约 (ArrowKingNft.sol)
- 基于ERC721Enumerable标准
- 支持批量铸造NFT
- 所有者权限控制

### 2. 拍卖合约 (AuctionContract.sol)
- 支持ETH和ERC20代币支付
- 集成Chainlink预言机价格转换
- 自动退还出价逻辑
- 拍卖结束自动转移NFT

### 3. 工厂合约 (AuctionFactoryContract.sol)
- 可升级的代理合约（UUPS模式）
- 批量创建拍卖合约
- 支持合约升级

## 🛠️ 技术栈

- **区块链框架**: Hardhat
- **智能合约**: Solidity ^0.8
- **NFT标准**: ERC721 (OpenZeppelin)
- **代理模式**: UUPS (Upgradeable)
- **预言机**: Chainlink
- **测试框架**: Chai + Mocha

## 📦 安装和部署

### 1. 安装依赖
```bash
npm install
```

### 2. 本地开发
```bash
# 部署合约
npx hardhat deploy --network localhost

# 运行测试
npx hardhat test test/AuctionTest.js
```

### 3. 测试网络部署
```bash
# 部署到Sepolia测试网
npx hardhat deploy --network sepolia
```

### 测试覆盖范围
- NFT铸造和查询
- 拍卖合约创建
- 买家参与拍卖
- 拍卖结束和NFT转移
- 工厂合约升级