// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract Credits is ERC20 {
    address private _owner;
    uint private _tokenPrice;

    event TokensRedeemed(address indexed redeemer, uint indexed amount);

    modifier onlyOwner {
        require(msg.sender == _owner, "CREDITS: only owner function");
        _;
    }
    constructor(string memory _name, string memory _symbol, uint tp) ERC20(_name, _symbol){
        _owner = msg.sender;
        _tokenPrice = tp;
    }


    function mintTokens(uint numOfTokens) public payable {
        require(msg.value >= _tokenPrice, "CREDITS: invalid ether amount sent");

        _mint(msg.sender, numOfTokens);

    }

    // owner mint to address function to use for individual airdrops
    function mintToAddress( address target, uint amount) public payable onlyOwner{
        _mint(target, amount);
    }

    // when users make redeem request, api request initiates the burn call from the owner wallet
    // tokens burned to redeem
    function redeemCredits( address target, uint amount) public onlyOwner(){
        _burn(target, amount);
        emit TokensRedeemed(target, amount);
    }

    // CREATE A HELPER FUNCTION TO UPDATE TOKEN PRICE WHENEVER MINTING FUNCTION IS CALLED




    // getter functions

    function owner() public view returns(address){
        return _owner;
    }
    function tokenPrice() public view returns(uint){
        return _tokenPrice;
    }

    function getContractBalance()public view onlyOwner returns(uint){
        return address(this).balance;
    }
}