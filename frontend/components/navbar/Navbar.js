// Navbar.js

import Link from 'next/link';
import WalletButton from '../walletbutton/WalletButton';
import styles from './navbar.module.css'; // Import the external CSS file
import { validateLogin, logout } from '../../actions/actions';
import Logout from '../logout/Logout';

const Navbar = async () => {
  const isValid = await validateLogin()

  return (
    <div className={styles.navbarContainer}>
      <div className={styles.linksContainer}>
        <Link href="/home">Home</Link>
        <Link href="/smartcontracts">Smart Contracts</Link>
        <Link href="/about">About</Link>
        <Link href="/getcredits">Get Credits</Link>
        {!isValid ? <Link href="/login">Login</Link> : <Logout /> }
        
      </div>
      <div>
        <WalletButton />
      </div>
    </div>
  );
};

export default Navbar;
