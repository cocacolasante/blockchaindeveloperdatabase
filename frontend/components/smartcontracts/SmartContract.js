import { getCookies, getSmartContract } from "../../actions/actions"


const SmartContract = async ({contractaddress}) => {
    const {apikey, address} = await getCookies()
    // @todo TEST SMART CONTRACT IS RETURNED PROPERLY
    const smartContract = await getSmartContract(contractaddress, apikey.value, address.value)
    
  return (
    <div>
      <h2>Smart contract</h2>
      <h1>Smart Contract Name: {smartContract.project_name}</h1>
      <p>Deployer Wallet: {smartContract.deployer_wallet}</p>
      <p>Target Address: {smartContract.address}</p>
      <p>Description: {smartContract.description}</p>
    </div>
  )
}

function getServerSideProps(context){

}

export default SmartContract