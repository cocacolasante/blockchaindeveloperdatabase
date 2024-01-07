require("@nomicfoundation/hardhat-toolbox");
require("dotenv").config()

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
  solidity: "0.8.20",
  networks:{
    polygon:{
      url: process.env.POLYGON_URL,
      accounts: [process.env.DEPLOYER_PRIVATE]
    },
    mumbai:{
      url: process.env.MUMBAI_URL,
      accounts: [process.env.DEPLOYER_PRIVATE]
    }
  }
};
