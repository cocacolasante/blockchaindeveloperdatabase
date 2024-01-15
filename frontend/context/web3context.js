"use client";
import React, {useState, useEffect, useContext} from "react";
import { ethers } from "ethers";
import { useRouter } from 'next/navigation'
import creditAbi from "../utils/abi/credits.json"
import {CREDITS_ADDRESS} from "../utils/addresses/addresses"
import { getCookie, setCookie } from 'cookies-next'
import Cookies from 'js-cookie';


// fetch smart contracts
const fetchCreditsContract = (signerOrProvider) =>{
    return new ethers.Contract(CREDITS_ADDRESS, creditAbi.abi, signerOrProvider)
}

export const SmartContractContext = React.createContext();

export const SmartContractProvider = ({children}) =>{
    
    
    const router = useRouter()
    const [currentAccount, setCurrentAccount] = useState()
    const [tokenPrice, setTokenPrice] = useState()
    const [loggedIn, setLoggedIn] = useState(false)

    const connectToWallet = async () =>{
        try{
            if(!window.ethereum){
                alert("please install metamask extension")
                
            }

            const accounts = await window.ethereum.request({method: "eth_requestAccounts"})
            if(accounts.length){
                setCurrentAccount(accounts[0]);
                await getTokenPrice()
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

                await getTokenPrice()
            }

        }catch(err){
            console.log(err )
        }
    }

    const getTokenPrice = async () => {
        try {
            if(!window.ethereum){
                return
            }
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
        if(tokenPrice == 0 ){
            console.log("tokenPrice equals 0")
            return
        }
    
        const purchaseAmount = (tokenamount * tokenPrice).toString()
        const sendValue = ethers.parseEther(purchaseAmount)

        const provider = new ethers.BrowserProvider(window.ethereum);
        const signer = await provider.getSigner()
        const contract = fetchCreditsContract(signer);

       
        const tx = await contract.mintTokens(tokenamount, {value: sendValue})
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
      
    }, [])
 


    return (
        <SmartContractContext.Provider
            value={({
                connectToWallet,
                checkIfWalletIsConnected,
                currentAccount,
                getTokenPrice,
                purchaseTokens,
                setTokenPrice,
                tokenPrice,
                checkIfWalletIsConnected,
                router
                

            })}
        >{children}</SmartContractContext.Provider>

    )
}
