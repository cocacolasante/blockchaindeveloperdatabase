const { expect } = require("chai");
const hre = require("hardhat")

describe("Credits Contract", () =>{
    let CreditContract, deployer, user1, user2
    beforeEach(async ()=>{
        [deployer, user1, user2] = await ethers.getSigners()
        CreditContract = await ethers.deployContract("Credits", ["Credit Token", "CT"])
        await CreditContract.waitForDeployment()

        console.log(`Contract Deployed to ${CreditContract.target}`)
        
    })
    describe("Deployment", () =>{
        it("checks owner address is deployer", async () =>{
            expect(await CreditContract.owner()).to.equal(deployer.address)
        })
        it("checks the ")

    })
})