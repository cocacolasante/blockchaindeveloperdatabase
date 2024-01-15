import { addContract } from "../../actions/actions";

const AddSmartcontract = () => {
  const addSC = async (e) => {
    "use server";
    e.preventDefault()

    const formData = new FormData(e.target);
    const name = formData.get("name");
    const address = formData.get("address");
    const deployer = formData.get("deployer");
    const description = formData.get("description");

    // Validate the form data and save it to the database

    console.log({ "project_name": name, "address": address, "deployer_wallet":deployer, "description":description });
    const response = await addContract(name, address, deployer, description)
    return response
  };

  // Form code

  return (
      <form method="POST" onSubmit={addSC} >
          <label htmlFor="contractaddress" >Address:</label>
          <input  id="address" type="text" name='address'/> 
          <label htmlFor="name" >Name:</label>
          <input  id="name" type="text" name='name' /> 
          <label htmlFor="name" >Deployer:</label>
          <input  id="deployer" type="text" name='deployer' /> 
          <label htmlFor="description" >description:</label>
          <input  id="description" type="text" name='description' /> 
          <button type="submit"  >Add Contract</button>
      </form>
  )
}

export default AddSmartcontract