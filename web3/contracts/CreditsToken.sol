// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract Credits is ERC20 {
    address public owner;

    uint public tokenPrice;

    event TokensRedeemed(address indexed redeemer, uint indexed amount);

    modifier onlyOwner {
        require(msg.sender == owner, "CREDITS: only owner function");
        _;
    }
    constructor(string memory _name, string memory _symbol) ERC20(_name, _symbol){
        owner = msg.sender;
    }


    function mintTokens(uint numOfTokens) public payable {
        require(msg.value >= tokenPrice, "CREDITS: invalid ether amount sent");

        _mint(msg.sender, numOfTokens);

    }

    // owner mint to address function to use for individual airdrops
    function mintToAddress(uint numOfTokens, address target) public payable onlyOwner{
        _mint(target, numOfTokens);
    }

    // when users make redeem request, api request initiates the burn call from the owner wallet
    // tokens burned to redeem
    function redeemCredits(uint amount, address target) public onlyOwner(){
        _burn(target, amount);
        emit TokensRedeemed(target, amount);
    }

    // CREATE A HELPER FUNCTION TO UPDATE TOKEN PRICE WHENEVER MINTING FUNCTION IS CALLED
}