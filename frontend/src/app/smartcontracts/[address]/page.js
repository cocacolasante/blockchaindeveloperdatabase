import React from 'react'
import SmartContract from '../../../../components/smartcontracts/SmartContract'

const page = (params) => {
  console.log(params.params.address)

  return (
    <div>
      <SmartContract contractaddress={params.params.address} />
    </div>
  )
}

export default page