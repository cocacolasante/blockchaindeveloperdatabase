"use client"
import { useState } from "react"
import { useContext, useEffect } from "react";
import { SmartContractContext } from "../../context/web3context"
import { useRouter } from "next/navigation";

const Signup = () => {

  const {currentAccount} = useContext(SmartContractContext);
  const [email, setEmail] = useState()
  const [password, setPassword] = useState()
  const [password2, setPassword2] = useState()

  const router = useRouter()
  const handleSignup = async (e) =>{
    e.preventDefault()
    if(!currentAccount){
      alert("Please connect MetaMask")
      return
    }
   
    if(password != password2){
      alert("Passwords do not match")
      return
    }

    const reqOptions = {
      method: "POST",
      headers: {
          'Content-Type': 'application/json',
      },
      body: JSON.stringify({ "email": email, "password": password, "wallet_address": currentAccount.toString()}),
  };

    const url = `http://localhost:8080/signup`

    const response = await fetch(url, reqOptions)
    const data = await response.json()
    console.log(data)
      
    alert("Please check email for activation link -then sign in again")
  }


  return (
    <div>
      <h2>Dont have an account? Sign up below</h2>
      <form method="POST"  onSubmit={handleSignup}>
        <label htmlFor="email" >Email:</label>
        <input onChange={e=>setEmail(e.target.value)} id="email" type="email" name='email'/> 
        <label htmlFor="password" >Password:</label>
        <input onChange={e=>setPassword(e.target.value)} id="password" type="password" name='password' /> 
        <label htmlFor="password2" >Re-Enter Password:</label>
        <input onChange={e=>setPassword2(e.target.value)}  id="password2" type="password" name='password2' /> 
        <button type="submit" onClick={handleSignup} >Sign Up</button>
      </form>
    </div>
  )
}

export default Signup