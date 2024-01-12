"use client"
import { useContext, useEffect } from "react";
import { SmartContractContext } from "../../context/web3context"

const WalletButton = () => {
    const {connectToWallet, currentAccount, checkIfWalletIsConnected} = useContext(SmartContractContext);

    useEffect(()=>{
      
      // checkIfLoggedIn()
  }, [currentAccount])

  return (
    <div>
    {!currentAccount ?  <button onClick={connectToWallet} >Connect To Wallet</button> : <p>{currentAccount}</p>}
    </div>
  )
}

export default WalletButton