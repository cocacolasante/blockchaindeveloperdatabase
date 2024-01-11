"use client"
import { useContext } from "react"
import { SmartContractContext } from "../../context/web3context"
import BuyCredits from "./buycredits"


const CreditsTop = () => {
    const {tokenPrice} = useContext(SmartContractContext)
    
  return (
    <div> 
         <h1>Purchase Credits</h1>
      <div>
        <h3>Price: {tokenPrice}</h3>
      </div>
      <div>
        <BuyCredits />
      </div>
    </div>
  )
}

export default CreditsTop