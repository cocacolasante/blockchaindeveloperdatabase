import { SmartContractProvider } from "../../context/web3context";
import Navbar from "../../components/navbar/Navbar";
import Footer from "../../components/footer/Footer";
import { validateLogin } from "../../actions/actions";
import Login from "../../components/login/Login";
import Signup from "../../components/signup/Signup";
import styles from "./global.css"

export const metadata = {
  title: "THE BLOCKCHAIN DEVELOPER DATABASE",
  description:
    "The best place to store blockchain developers to store their deployed smart contract information",
};

export default async function RootLayout({ children }) {
  const valid = await getValidation()
  
  if (!valid ){
    return (
    <html lang="en">
        <body className={styles.body}>
          <SmartContractProvider>
          <Navbar />
            <Login />

            <Signup />
          </SmartContractProvider>
        </body>
      </html>)
  }
  return (
    <html lang="en">
      <SmartContractProvider>
        <body>
          <Navbar />
          {children}
          <Footer />
        </body>
      </SmartContractProvider>
    </html>
  );
}

export async function getValidation() {
  const isValidated = await validateLogin();

  return isValidated
}
