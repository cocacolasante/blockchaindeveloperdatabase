"use client"
import { useState, useContext } from "react"
import { SmartContractContext } from "../../context/web3context"
import { redirect } from "next/dist/server/api-utils";

const Logout = () => {
    const {router} = useContext(SmartContractContext);
    const handleLogout = async () =>{
        const options = {
            method: 'GET',
            credentials: 'include',
            headers: {
            'Content-Type': 'application/json',
            }
           
        };

        let response = await fetch("http://localhost:8080/logout", options)
        
        const data = await response.json()
        
        console.log(data)
        if(data["logged_out"]){
            router.refresh()
        }
        
    }
  return (
    
    <a onClick={handleLogout}>
        Logout
    </a>
  )
}

export default Logout