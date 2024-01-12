"use client"
import { useCookies } from 'next-client-cookies';

export async function logout() {
    
    const cookieStore = useCookies();
    
    cookieStore.set("apikey", "")
    cookieStore.set("email", "")

    console.log(cookieStore.get("apikey"))
    console.log(cookieStore.get("email"))

}