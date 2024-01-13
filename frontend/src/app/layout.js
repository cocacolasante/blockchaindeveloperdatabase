import { SmartContractProvider } from "../../context/web3context";
import Navbar from "../../components/navbar/Navbar";
import Footer from "../../components/footer/Footer";
import { validateLogin } from "../../actions/actions";
import Login from "../../components/login/Login";

export const metadata = {
  title: "THE BLOCKCHAIN DEVELOPER DATABASE",
  description:
    "The best place to store blockchain developers to store their deployed smart contract information",
};

export default async function RootLayout({ children }) {
  const valid = await getValidation()
  console.log(valid)
  if (!valid ){

    return (
    <html lang="en">
        <body>
          <Login />
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
