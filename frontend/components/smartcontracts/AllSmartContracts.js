import { getAllSmartContracts } from "../../actions/actions"
import SmartContractList from "../smartcontractcomponents/SmartContractList"


const AllSmartContracts = async () => {
    const smartcontracts = await getAllSmartContracts()

    if(smartcontracts.length == 0 ){
      return(
        <div>
        <h2>No smart contracts in database</h2>
      </div>
      )
    }
    
  return (
    <div>
      <h1>All Smart Contracts</h1>
      
        {!smartcontracts && smartcontracts.map((contract, i) =>{
          return (
            <div key={i}>
              <SmartContractList address={contract.address} description={contract.description} name={contract.project_name} />
            </div>
          )
        })}
      
    </div>
  )
}

export default AllSmartContracts