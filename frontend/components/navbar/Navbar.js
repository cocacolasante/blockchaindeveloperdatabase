// Navbar.js

import Link from 'next/link';
import WalletButton from '../walletbutton/WalletButton';
import styles from './navbar.module.css'; // Import the external CSS file

const Navbar = () => {
  return (
    <div className={styles.navbarContainer}>
      <div className={styles.linksContainer}>
        <Link href="/home">Home</Link>
        <Link href="/smartcontracts">Smart Contracts</Link>
        <Link href="/about">About</Link>
        <Link href="/login">Login</Link>
        <Link href="/getcredits">Get Credits</Link>
      </div>
      <div>
        <WalletButton />
      </div>
    </div>
  );
};

export default Navbar;
