const { expect } = require("chai");
const hre = require("hardhat")

describe("Credits Contract", () =>{
    let CreditContract, deployer, user1, user2
    beforeEach(async ()=>{
        [deployer, user1, user2] = await ethers.getSigners()
        CreditContract = await ethers.deployContract("Credits", ["Credit Token", "CT", ethers.parseUnits("1")])
        await CreditContract.waitForDeployment()

        // console.log(`Contract Deployed to ${CreditContract.target}`)
        
    })
    describe("Deployment", () =>{
        it("checks owner address is deployer", async () =>{
            expect(await CreditContract.owner()).to.equal(deployer.address)
        })
        it("checks the token price is 1 ether", async () =>{
            expect(await CreditContract.tokenPrice()).to.equal(ethers.parseUnits("1"))
        })

    })
    describe("Token Minting: User", () =>{
        describe("Success", () =>{
            beforeEach(async () =>{ 
                await CreditContract.connect(user1).mintTokens(10,{value: ethers.parseUnits("10")})
                
            })
            it("checks user1 balance", async () =>{
                expect(await CreditContract.balanceOf(user1.address)).to.equal(10)
            })
            it("checks the contract balance to equal 10 eth", async () =>{
                // TO DO ------ CHECK FOR FAILURE WHEN CALLED
                expect(await CreditContract.getContractBalance()).to.equal(ethers.parseUnits("10"))
            })
        })
        describe("Failure", () =>{
            it("expects fail for not enough ether sent", async () =>{
                await expect(CreditContract.mintTokens(10)).to.be.reverted
            })
            it("expects call to fail for owner not calling it", async () =>{
                await expect(CreditContract.connect(user1).getContractBalance()).to.be.reverted
            })
        })
    })
    describe("Minting Tokens - Admin",() =>{
        describe("Success",() =>{
            beforeEach(async () =>{
                await CreditContract.connect(deployer).mintToAddress(user1.address, 10)
            })
            it("expects user 1 balance to equal 10 tokens to redeem", async () =>{
                expect(await CreditContract.balanceOf(user1.address)).to.equal(10)
            })
        })
    })
})