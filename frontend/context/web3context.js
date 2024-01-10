"use client";
import React, {useState, useEffect, useContext} from "react";
import { ethers } from "ethers";



// import { ProfileAbi, ProfileContractAddress, AlbumCreatorAbi, AlbumCreatorAddress, networkId, rpcConnection, AlbumNftAbi } from "./constants";

// fetch smart contracts
// const fetchProfileContract = (signerOrProvider) =>{
//     return new ethers.Contract(ProfileContractAddress, ProfileAbi, signerOrProvider)
// }
// const fetchAlbumContract = (signerOrProvider) =>{
//     return new ethers.Contract(AlbumCreatorAddress, AlbumCreatorAbi, signerOrProvider)
// }




export const SmartContractContext = React.createContext();

export const SmartContractProvider = ({children}) =>{
    const [currentAccount, setCurrentAccount] = useState()

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


    const RpcProvider = new ethers.BrowserProvider()

    useEffect(()=>{
        checkIfWalletIsConnected();
    }, [])

    return (
        <SmartContractContext.Provider
        value={({
            connectToWallet,
            checkIfWalletIsConnected,
            currentAccount,
            RpcProvider,

        })}
        >{children}</SmartContractContext.Provider>

    )
}
