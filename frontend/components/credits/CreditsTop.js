"use client"
import React, { useContext, useEffect } from "react";
import { SmartContractContext } from "../../context/web3context";
import BuyCredits from "./BuyCredits";
import styles from "./CreditTop.module.css"
import UsersBalance from "./UsersBalance";

const CreditsTop = () => {
  const { getTokenPrice, tokenPrice, setTokenPrice, currentAccount } = useContext(SmartContractContext);



  return (
    <div className={styles.creditsTopContainer}>
      <h1 className={styles.heading}>Purchase Credits</h1>
      <div>
        <h3 className={styles.subHeading}>Price: {tokenPrice} Matic</h3>
      </div>
      <div>
        <BuyCredits />
      </div>
    </div>
  );
};

export default CreditsTop;
