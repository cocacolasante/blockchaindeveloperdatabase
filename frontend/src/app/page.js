import { validateLogin } from "../../actions/actions"


export default function Home() {
  return (
    <div>
     {console.log(validateLogin())}
      <h1>THE BLOCKCHAIN DEVELOPER DATABASE</h1>
    </div>
  )
}
