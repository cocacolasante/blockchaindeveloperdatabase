"use client"
import { useState, useContext } from "react"
import { SmartContractContext } from "../../context/web3context"
import { useRouter } from "next/navigation";
import Link from "next/link";

const Login = () => {
    const router = useRouter()
    const [emailInput, setEmailInput] = useState("")
    const [passwordInput, setPasswordInput] = useState("")

    const handleLogin = async (e) =>{
        e.preventDefault()
        if(emailInput == ""){
            alert("enter email")
            return
        }
        if(passwordInput == ""){
            alert("enter password")
            return
        }
        try {
            const email = emailInput
            const password = passwordInput

            // Construct the options for the fetch request
            const options = {
                method: 'POST',
                credentials: 'include',
                headers: {
                'Content-Type': 'application/json',
                },
                body: JSON.stringify({ "email":email, "password": password }),
            };

            let response = await fetch("http://localhost:8080/login", options)
            
            const data = await response.json()
            
            if(data["api_key"]){
                router.push("/smartcontracts")
                router.refresh()
            }else {
                alert("login failed")
            }
        }catch(err){
            console.log(err)
            return
        }
    }


  return (
    <div>
        <h1>THE BLOCKCHAIN DATABASE</h1>
        <h4>Please login</h4>
   
    <form onSubmit={handleLogin}>
        <label htmlFor="email" >Email:</label>
        <input onChange={e=>setEmailInput(e.target.value)} id="email" type="email" name='email'/> 
        <label htmlFor="password" >Password:</label>
        <input onChange={e=>setPasswordInput(e.target.value)} id="password" type="password" name='password' /> 
        <button type="submit" onClick={handleLogin} >Login</button>
        <Link href={`/signup`}>Signup</Link>
    </form>
    </div>
  )
}

export default Login