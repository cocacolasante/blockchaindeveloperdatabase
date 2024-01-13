import { cookies } from 'next/headers';
import { getWalletFromState } from './clientactions';

export async function validateLogin() {
  
    const cookieStore = cookies();
    const apikey = cookieStore.get("apikey");
    const email = cookieStore.get("email");
    

    if(!apikey){
        console.log("apikey does not exist")
        return
    }
    if(!email){
        console.log("email")
        return
    }

    const isLoggedIn = await checkApikeyToDatabase(apikey, email);
    console.log(isLoggedIn);
    return isLoggedIn;
}

export async function checkApikeyToDatabase(apikey, email) {
    if(!apikey || !email ){
        return
    }
    console.log(apikey)
    console.log(email)
    const reqOptions = {
        method: "POST",
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ "email": email.value, "api_key": apikey.value }),
    };

    const response = await fetch("http://localhost:8080/validatekey", reqOptions);
    const data = await response.json();
    

    return data.matches; // Assuming the server sends a boolean field named "matches"
}


export async function getSmartContract(contractAddress){
    const reqOptions = {
        method: "GET",
        headers: {
            'Content-Type': 'application/json',
        }      
    };
    const useraddress = getWalletFromState()
    console.log(useraddress)



    const fetchUrl = `localhost:8080/${useraddress}/contract/${contractAddress}`
    const response = await fetch(fetchUrl, reqOptions)
    const data = await response.json()

    return data

}

export async function getCookies() {
    const cookieStore = cookies();
    const apikey = cookieStore.get("apikey");
    const email = cookieStore.get("email");

    return {
        email,
        apikey
    }
}