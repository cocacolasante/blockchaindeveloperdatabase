"use client"
import { useContext, useState } from "react"
import { SmartContractContext } from "../../context/web3context"

const BuyCredits = () => {
    const [tokensAmount, setTokensAmount] = useState()
    const {purchaseTokens} = useContext(SmartContractContext)

    const handlePurchase = async () =>{
        try {
            if (tokensAmount == 0 ){
                console.log("error zero tokens")
                return
            }

            
            purchaseTokens(tokensAmount)

        }catch(err){
            console.log(err)
        }
    }

  return (
    <div>
        <label htmlFor="tokenamount" >Amount of Tokens To Purchase</label>
        <input onChange={e=>setTokensAmount(e.target.value)} name="tokenamount" type="number"></input>
        <button onClick={handlePurchase} >Purchase</button>
    </div>
  )
}

export default BuyCredits