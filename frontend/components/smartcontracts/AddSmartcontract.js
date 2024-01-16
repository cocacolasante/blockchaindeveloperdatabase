import { addContract } from "../../actions/actions";


const AddSmartcontract = () => {

  return (
      <form method="POST" action={addContract} >
          <label htmlFor="contractaddress" >Address:</label>
          <input  id="address" type="text" name='address'/> 
          <label htmlFor="project_name" >Name:</label>
          <input  id="project_name" type="text" name='project_name' /> 
          <label htmlFor="deployer_wallet" >Deployer:</label>
          <input  id="deployer_wallet" type="text" name='deployer_wallet' /> 
          <label htmlFor="description" >description:</label>
          <input  id="description" type="text" name='description' /> 
          <button type="submit"  >Add Contract</button>
      </form>
  )
}

export default AddSmartcontract