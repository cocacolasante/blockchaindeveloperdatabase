import React from 'react'
import CreditsTop from './CreditsTop'
import styles from "./CreditTop.module.css"
import UsersBalance from './UsersBalance'

const CreditsContainer = () => {
  return (
    <div className={styles.creditsTopContainer}>
        <UsersBalance  />
        <CreditsTop />
  </div>

  )
}

export default CreditsContainer