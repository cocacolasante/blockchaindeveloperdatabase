import { ethers } from "ethers";
import { cookies } from 'next/headers'
import { useRouter } from 'next/navigation'
import creditAbi from "../utils/abi/credits.json"
import {CREDITS_ADDRESS} from "../utils/addresses/addresses"
import { useEffect } from "react";
import { validateLogin } from "../actions/actions";


export const LoginContext = React.createContext();

export const LoginProvider = ({children}) =>{
   
    
    
    
    return (
        <LoginContext.Provider
            value={({
            })}
        >{children}</LoginContext.Provider>

    )
}
