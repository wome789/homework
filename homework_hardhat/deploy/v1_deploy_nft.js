const { ethers } = require("hardhat");

module.exports = async function ({ getNamedAccounts, deployments }) {
    console.log("=================================部署nft合约开始=================================");
    const { save } = deployments;

    const { deployer } = await getNamedAccounts();

    // 部署NFT合约
    const nftContractFactory = await ethers.getContractFactory("ArrowKingNft");
    const nftContract = await nftContractFactory.deploy();
    await nftContract.waitForDeployment();
    const nftContractAddress = await nftContract.getAddress();
    console.log("NFTContract deployed at: ", nftContractAddress);

    // 写入部署记录
    await save("depolyNft", {
        abi: nftContractFactory.interface.format("json"),
        address: nftContractAddress,
        from: deployer,
        args: []
    });
    console.log("=================================部署nft合约完成=================================");
}
module.exports.tags = ["depolyNft"];