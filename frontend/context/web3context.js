"use client";
import React, {useState, useEffect, useContext} from "react";
import { ethers } from "ethers";

import creditAbi from "../utils/abi/credits.json"
import {CREDITS_ADDRESS} from "../utils/addresses/addresses"

// fetch smart contracts
const fetchCreditsContract = (signerOrProvider) =>{
    return new ethers.Contract(CREDITS_ADDRESS, creditAbi.abi, signerOrProvider)
}

export const SmartContractContext = React.createContext();

export const SmartContractProvider = ({children}) =>{
    const [currentAccount, setCurrentAccount] = useState()
    const [tokenPrice, setTokenPrice] = useState()

    const connectToWallet = async () =>{
        try{
            if(!window.ethereum){
                alert("please install metamask extension")
                
            }

            const accounts = await window.ethereum.request({method: "eth_requestAccounts"})
            if(accounts.length){
                setCurrentAccount(accounts[0]);

            }

        }catch(err){
            console.log(err )
        }
    }

    const checkIfWalletIsConnected = async () =>{
        try{
            if(!window.ethereum){
                alert("please install metamask extension")
            }

            const accounts = await window.ethereum.request({method: "eth_accounts"})
            if(accounts.length){
                const currentUser = accounts[0];
    
                setCurrentAccount(currentUser);

            }

        }catch(err){
            console.log(err )
        }
    }

    const getTokenPrice = async () => {
        try {
          const provider = new ethers.BrowserProvider(window.ethereum);
          const contract = fetchCreditsContract(provider);
      
          const creditPrice = await contract.tokenPrice();
          
            const parsed = ethers.formatEther(creditPrice)
          setTokenPrice(parsed);
        } catch (error) {
          console.error("Error fetching token price:", error);
        }
      };
      
      const purchaseTokens = async (tokenamount) =>{
          try {
            const purchaseAmount = (tokenamount * tokenPrice).toString()


            const provider = new ethers.BrowserProvider(window.ethereum);
            const signer = await provider.getSigner()
            const contract = fetchCreditsContract(signer);

            const tx = await contract.mintTokens(1, {value: purchaseAmount})
            const res = await tx.wait()
            if(res.status == 1) {
                console.log("Success")
            }else {
                console.log("Failed")
                console.log("Transaction failed. Revert reason:", receipt.events[0]?.args?.message);
            }
        
        }catch(err){
            console.log(err)
        }
      }

    useEffect(()=>{
        checkIfWalletIsConnected();
        getTokenPrice()
       
    }, [])


    return (
        <SmartContractContext.Provider
            value={({
                connectToWallet,
                checkIfWalletIsConnected,
                currentAccount,
                tokenPrice,
                purchaseTokens

            })}
        >{children}</SmartContractContext.Provider>

    )
}
