
import { SmartContractProvider } from "../../context/web3context"
import Navbar from '../../components/navbar/Navbar'
import Footer from "../../components/footer/Footer"
import { validateLogin } from "../../actions/actions"
import Login from "../../components/login/Login"


export const metadata = {
  title: 'THE BLOCKCHAIN DEVELOPER DATABASE',
  description: 'The best place to store blockchain developers to store their deployed smart contract information',
}




export default async function RootLayout({ children }) {
  const isValidated = await validateLogin()
  
  return (
    <html lang="en">
      <SmartContractProvider>
          <body >
            <Navbar /> 
        {isValidated ? (
          <>

            {children}
           
          </>
         
        ) : (
          <>
             
              <Login /> 
             
          </>
        )}
            <Footer />
          </body>
        
      </SmartContractProvider>
    </html>
  )
}
