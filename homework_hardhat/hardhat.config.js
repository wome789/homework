require("@nomicfoundation/hardhat-toolbox");
require('hardhat-deploy');
require("@openzeppelin/hardhat-upgrades")

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
  solidity: "0.8.28",
  namedAccounts: {
    deployer: 0,
    buyer1: 1,
    buyer2: 2
  },
  networks: {
    sepolia: {
      url: `https://sepolia.infura.io/v3/${process.env.INFURA_API_KEY}`,
      accounts: [process.env.PK]
    }
  }
};
