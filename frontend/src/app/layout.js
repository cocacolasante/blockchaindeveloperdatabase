
import { SmartContractProvider } from "../../context/web3context"
import Navbar from '../../components/navbar/Navbar'
import Footer from "../../components/footer/Footer"
import { CookiesProvider } from 'next-client-cookies/server';
import { LoginProvider } from "../../context/logincontext";

export const metadata = {
  title: 'THE BLOCKCHAIN DEVELOPER DATABASE',
  description: 'The best place to store blockchain developers to store their deployed smart contract information',
}

export default function RootLayout({ children }) {
  return (
    <html lang="en">
        <SmartContractProvider>
            <body >
              <Navbar />
                {children}
              <Footer />
            </body>
        </SmartContractProvider>
    </html>
  )
}
