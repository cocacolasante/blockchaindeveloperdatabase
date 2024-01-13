"use client"
import { useContext, useEffect } from "react";
import { SmartContractContext } from "../../context/web3context"

export async function getWalletFromState() {
    const {currentAccount} = useContext(SmartContractContext);


    return currentAccount


}