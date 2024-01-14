"use client"
import { useState } from "react"
import { useContext, useEffect } from "react";
import { SmartContractContext } from "../../context/web3context"



const AddSmartcontract = () => {
  const {currentAccount} = useContext(SmartContractContext);

  const [address, setAddress] = useState()
  const [deployer, setDeployer] = useState()
  const [name, setName] = useState()
  const [description, setDescription ] = useState()

  const handleAddContract = async (e) =>{
    e.preventDefault()
    if(!address || !deployer || !name || !description){
      return
    }

    try{ 
      // @todo create fetch to backend to post add smart contract form
      const wallet = await fetch(`http://localhost:8080/${currentAccount}/`)
      // console.log(data)

    }catch(err){
      console.log(err)
    }

  }

  return (
    <div>
      <form>
        <label htmlFor="contractaddress" >Address:</label>
        <input onChange={e=>setAddress(e.target.value)} id="address" type="text" name='address'/> 
        <label htmlFor="name" >Name:</label>
        <input onChange={e=>setName(e.target.value)} id="name" type="text" name='name' /> 
        <label htmlFor="name" >Deployer:</label>
        <input onChange={e=>setDeployer(e.target.value)} id="deployer" type="text" name='deployer' /> 
        <label htmlFor="description" >description:</label>
        <input onChange={e=>setDescription(e.target.value)} id="description" type="text" name='description' /> 
        <button type="submit" onClick={handleAddContract} >Login</button>
      </form>
    </div>
  )
}

export default AddSmartcontract