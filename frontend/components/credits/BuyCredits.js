"use client"
// BuyCredits.jsx
import React, { useContext, useState } from "react";
import { SmartContractContext } from "../../context/web3context";
import styles from "./BuyCredits.module.css";

const BuyCredits = () => {
  const [tokensAmount, setTokensAmount] = useState();
  const { purchaseTokens } = useContext(SmartContractContext);

  const handlePurchase = async () => {
    try {
      if (tokensAmount == 0) {
        console.log("error zero tokens");
        return;
      }
      await purchaseTokens(tokensAmount);
    } catch (err) {
      console.log(err);
    }
  };

  return (
    <div className={styles.buyCreditsContainer}>
    
      <label htmlFor="tokenamount" className={styles.label}>
        Amount of Tokens To Purchase
      </label>
      <input
        onChange={(e) => setTokensAmount(e.target.value)}
        name="tokenamount"
        type="number"
        className={styles.input}
      />
      <button onClick={handlePurchase} className={styles.button}>
        Purchase
      </button>
    </div>
  );
};

export default BuyCredits;
