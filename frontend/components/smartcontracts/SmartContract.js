import { getSmartContract } from "../../actions/actions"


const SmartContract = async () => {
    // @todo TEST SMART CONTRACT IS RETURNED PROPERLY
    const smartContract = await getSmartContract()

  return (
    <div>
      <h1>Smart Contract Name: {smartContract.project_name}</h1>
    </div>
  )
}

export default SmartContract