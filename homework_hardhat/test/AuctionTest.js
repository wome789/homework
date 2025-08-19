const { ethers, deployments, upgrades } = require("hardhat");
const { expect } = require("chai");

describe("测试所有拍卖合约相关事项 Auction", function () {

    let nftContract, auctionFactoryContractProxy, auctionContract;
    let depolyNftInfo;
    let deployer, user1, user2;
    const tokenId = 1;

    let depolyAuctionFactoryContractProxyInfo;

    // 部署合约
    this.beforeEach(async function () {
        [deployer, user1, user2] = await ethers.getSigners();

        // 部署nft合约
        await deployments.fixture(["depolyNft"]);
        // 获取nft合约数据
        depolyNftInfo = await deployments.get("depolyNft");
        // 获取nft合约实例
        nftContract = await ethers.getContractAt("ArrowKingNft", depolyNftInfo.address);

        // 部署工厂合约
        await deployments.fixture(["AuctionFactoryContractProxy"]);
        // 获取工厂合约数据
        depolyAuctionFactoryContractProxyInfo = await deployments.get("AuctionFactoryContractProxy");
        // 获取代理工厂合约实例
        auctionFactoryContractProxy = await ethers.getContractAt("AuctionFactoryContract", depolyAuctionFactoryContractProxyInfo.address);
    })

    it("全量测试", async function () {
        // 创建10个nft
        for (let i = 0; i < 10; i++) {
            await nftContract.mint(deployer.address, i + 1);
        }
        console.log("测试创建Nft完成");

        // 获取nft信息
        const tokenIds = await nftContract.getTokenIds();
        console.log("测试获取nft信息完成", tokenIds);


        // 使用工厂合约创建拍卖合约
        let currentSeconds = Math.floor(Date.now() / 1000);
        const _startPrice = 1;
        const _startTime = currentSeconds;
        const _endTime = currentSeconds + 100;
        const _nftContract = depolyNftInfo.address;

        const auctionAddress = await auctionFactoryContractProxy.createAuction(_startPrice, _startTime, _endTime, _nftContract, tokenId);
        await auctionAddress.wait();


        // 获取拍卖合约信息并授权NFT
        let auctionContractAddress = await auctionFactoryContractProxy.getAuction(tokenId);
        // 授权nft给部署者
        await nftContract.connect(deployer).setApprovalForAll(auctionContractAddress, true);
        // 获取拍卖合约实例
        auctionContract = await ethers.getContractAt("AuctionContract", auctionContractAddress);

        // 使用ETH参与拍卖，传入金额和货币地址(0x0表示ETH)
        await auctionContract.connect(user1).bid(0, ethers.ZeroAddress, {
            value: ethers.parseEther("1")
        });
        console.log("测试买家参与拍卖完成");


        const auctionInfo = await auctionContract.getAuction();
        console.log("测试获取拍卖合约信息完成2", auctionInfo);

        // 拍卖结束
        /* await new Promise((resolve, reject) => {
            setTimeout(resolve, 4000);
        }) */
        await auctionContract.endAuction();
    })

    it("测试工厂合约升级", async function () {
        // 部署V2版本合约信息
        const AuctionFactoryContractV2 = await ethers.getContractFactory("AuctionFactoryContractV2");

        const address = await auctionFactoryContractProxy.getAddress();
        const implementAddress = await upgrades.erc1967.getImplementationAddress(address);
        console.log("升级前的合约地址", address);
        console.log("升级前的实现合约地址", implementAddress);

        // 升级合约
        auctionFactoryContractProxy = await upgrades.upgradeProxy(depolyAuctionFactoryContractProxyInfo.address, AuctionFactoryContractV2);
        await auctionFactoryContractProxy.waitForDeployment();

        const address2 = await auctionFactoryContractProxy.getAddress();
        const implementAddress2 = await upgrades.erc1967.getImplementationAddress(address2);
        console.log("升级后的合约地址", address2);
        console.log("升级后的实现合约地址", implementAddress2);

        // 调用升级后的方法
        const result = await auctionFactoryContractProxy.testUpgrade();
        expect(result).to.equal("test");

    })

})