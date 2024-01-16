import { getUsersBalance } from "../../actions/actions"


const UsersBalance = async () => {
    const userBalance = await getUsersBalance()

  return (
    <div>
        {userBalance &&  <h3>Users Current Balance: {userBalance.credits_available}</h3>}
       
    </div>
  )
}

export default UsersBalance