"use client"
import { useState, useContext } from "react"
import { SmartContractContext } from "../../context/web3context"

const Login = () => {
    const {router} = useContext(SmartContractContext);
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
            console.log(response)
            const data = await response.json()
            console.log(data["api_key"])
            if(data["api_key"]){
                router.push("/home")
            }else {
                alert("login failed")
            }
        }catch(err){
            console.log(err)
            return
        }
    }


  return (
    <form onSubmit={handleLogin}>
        <label htmlFor="email" >Email:</label>
        <input onChange={e=>setEmailInput(e.target.value)} id="email" type="email" name='email'/> 
        <label htmlFor="password" >Password:</label>
        <input onChange={e=>setPasswordInput(e.target.value)} id="password" type="password" name='password' /> 
        <button type="submit" onClick={handleLogin} >Login</button>
        <button >Signup</button>
    </form>
  )
}

export default Login