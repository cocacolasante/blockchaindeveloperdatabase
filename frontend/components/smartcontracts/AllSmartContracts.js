import Link from "next/link"
import { getAllSmartContracts } from "../../actions/actions"
import SmartContractList from "../smartcontractcomponents/SmartContractList"


const AllSmartContracts = async () => {
    const smartcontracts = await getAllSmartContracts()
    
    if(smartcontracts.error){
      return(
        <div>
        <h2>No smart contracts in database</h2>
        <div>
        <Link href={"/smartcontracts/addcontract"} >Add Smart Contract </Link>

      </div>
      </div>
      )
    }
    
  return (
    <div>
      <div>
        <h1>All Smart Contracts</h1>
      </div>
      <div>
        <Link href={"/smartcontracts/addcontract"} >Add Smart Contract </Link>

      </div>
      {!smartcontracts ? (
        <div>
        <h2>No smart contracts in database</h2>
      </div>
      ): (
        smartcontracts.map((contract, i) =>{
          return (
            <div key={i}>
              <SmartContractList address={contract.address} description={contract.description} name={contract.project_name} />
            </div>
          )
        }
      ))}
        
      
    </div>
  )
}

export default AllSmartContracts