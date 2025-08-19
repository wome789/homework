const { ethers, upgrades, deployments } = require("hardhat");

module.exports = async function ({ getNamedAccounts, deployments }) {
    console.log("=================================部署工厂合约开始=================================");

    const { save } = deployments;
    const { deployer } = await getNamedAccounts();

    // 部署工厂合约的代理合约
    const auctionFactoryContractFactory = await ethers.getContractFactory("AuctionFactoryContract");
    // 创建代理合约
    const auctionFactoryContractProxy = await upgrades.deployProxy(auctionFactoryContractFactory, [], {
        initializer: "initialize"
    })

    // 获取代理工厂合约地址
    const auctionFactoryContractProxyAddress = await auctionFactoryContractProxy.getAddress();
    console.log("AuctionFactoryContract deployed at: ", auctionFactoryContractProxyAddress);

    // 获取工厂合约的逻辑合约地址
    const auctionFactoryContractProxyImplementationAddress = await upgrades.erc1967.getImplementationAddress(auctionFactoryContractProxyAddress);
    console.log("AuctionFactoryContract implementation deployed at: ", auctionFactoryContractProxyImplementationAddress);

    // 写入部署记录
    await save("AuctionFactoryContractProxy", {
        abi: auctionFactoryContractFactory.interface.format("json"),
        address: auctionFactoryContractProxyAddress,
        from: deployer,
        args: []
    })

    console.log("=================================部署工厂合约完成=================================");
}
module.exports.tags = ["AuctionFactoryContractProxy"]