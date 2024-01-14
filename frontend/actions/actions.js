import { cookies } from 'next/headers';

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


export async function getSmartContract(contractAddress, apikey, userAddress){

    const reqOptions = {
        method: "GET",
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${apikey}`
        }      
    };
   
    

// TEST URL
// http://localhost:8080/0x11273F391609BF4C05CA23c6aD29D919a71dc37E/contract/0xD9A9431cA0FbC045B96F66fF01762D63E7E113be

    const fetchUrl = `http://localhost:8080/${userAddress}/contract/${contractAddress}`
    
    const response = await fetch(fetchUrl, reqOptions)

    const data = await response.json()
    // console.log(data)
    return data

}

export async function getCookies() {
    const cookieStore = cookies();
    const apikey = cookieStore.get("apikey");
    const email = cookieStore.get("email");
    const address = cookieStore.get("useraddress");

    return {
        email,
        apikey,
        address
    }
}


export async function getAllSmartContracts(){
    const {email, apikey, address} = await getCookies()
    const reqOptions = {
        method: "GET",
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${apikey.value}`
        }      
    };
   
    const url = `http://localhost:8080/${address.value}/fullcontracts`

    const response = await fetch(url, reqOptions)
    const data = await response.json()
    return data
}




export async function addContract(name, contractAddress, deployer, description) {
    const {email, apikey, address} = await getCookies()

    const reqOptions = {
        method: "POST",
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${apikey.value}`
        },
        body: JSON.stringify({ "project_name": name, "address": contractAddress, "deployer_wallet": deployer, "description": description }),
    };

    const url = `http://localhost:8080/${address}/newcontract`
    const response = await fetch(url, reqOptions)
    const data = await response.json()
    return data
}