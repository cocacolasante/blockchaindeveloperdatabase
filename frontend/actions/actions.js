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

