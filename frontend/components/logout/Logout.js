"use client"
import { useState, useContext } from "react"
import { SmartContractContext } from "../../context/web3context"

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
        if(data.Logout){
            router.reload()
        }
        
    }
  return (
    
    <a onClick={handleLogout}>
        Logout
    </a>
  )
}

export default Logout