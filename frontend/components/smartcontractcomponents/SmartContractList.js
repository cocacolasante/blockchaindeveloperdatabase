import Link from 'next/link'
import React from 'react'

const SmartContractList = ({address, name, description}) => {
  return (
    <div >
        <Link href={`/smartcontracts/${address}`}>
            <p>Project Name: {name}</p>
            <p>Address: {address}</p>
            <p>Description: {description}</p>

        </Link>
    </div>
  )
}

export default SmartContractList