// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

interface Credits {
    event TokensRedeemed(address indexed redeemer, uint indexed amount);

    function mintTokens(uint numOfTokens) external payable;

    // owner mint to address function to use for individual airdrops
    function mintToAddress( address target, uint amount) external payable;

    // when users make redeem request, api request initiates the burn call from the owner wallet
    // tokens burned to redeem
    function redeemCredits( address target, uint amount) external;


    function withdrawl() external;

    // getter functions
    function owner() external view returns(address);
    function tokenPrice() external view returns(uint);

    function getContractBalance()external view returns(uint);

    function usersRedeemed(address user) external view returns(uint);
}