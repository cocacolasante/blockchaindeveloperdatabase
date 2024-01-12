"use client";
import React, {useState, useEffect, useContext} from "react";
import { ethers } from "ethers";
import { useCookies } from 'next-client-cookies';
import { useRouter } from 'next/navigation'
import creditAbi from "../utils/abi/credits.json"
import {CREDITS_ADDRESS} from "../utils/addresses/addresses"


export const LoginContext = React.createContext();

export const LoginProvider = ({children}) =>{
    const cookieStore = useCookies()
    
    const router = useRouter()

    const checkIfLoggedIn = () =>{
        const auth = cookieStore.get("api_key")
        if(!auth){
            router.push("/login")
        }

    }

    useEffect(()=>{

        checkIfLoggedIn()
    }, [])

  


    return (
        <LoginContext.Provider
            value={({
            })}
        >{children}</LoginContext.Provider>

    )
}
