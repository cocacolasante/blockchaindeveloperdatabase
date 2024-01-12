import { ethers } from "ethers";
import { cookies } from 'next/headers'
import { useRouter } from 'next/navigation'
import creditAbi from "../utils/abi/credits.json"
import {CREDITS_ADDRESS} from "../utils/addresses/addresses"


export const LoginContext = React.createContext();

export const LoginProvider = ({children}) =>{
    const cookieStore = cookies()
    
    const router = useRouter()

    const checkIfLoggedIn = () =>{
        const auth = cookieStore.get("apiKey")
        if(!auth){
            router.push("/login")
        }

    }
    return (
        <LoginContext.Provider
            value={({
            })}
        >{children}</LoginContext.Provider>

    )
}
