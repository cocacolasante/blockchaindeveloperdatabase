"use client"
import { useContext } from "react";
import { SmartContractContext } from "../../context/web3context"

const WalletButton = () => {
    const {connectToWallet, currentAccount} = useContext(SmartContractContext);
  return (
    <div>
    {!currentAccount ?  <button onClick={connectToWallet} >Connect To Wallet</button> : <p>{currentAccount}</p>}
       {console.log(currentAccount)}
    </div>
  )
}

export default WalletButton