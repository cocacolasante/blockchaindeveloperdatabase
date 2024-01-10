-- Create TABLE walletaccounts
CREATE TABLE walletaccounts (
    wallet_address VARCHAR(255) PRIMARY KEY,
    email VARCHAR(255) UNIQUE,
    password VARCHAR(255),
    smart_contract_addresses VARCHAR(255)[], -- Assuming array of strings can be stored as JSON
    credits_available BIGINT,
    api_key VARCHAR(255),
    activated BOOLEAN DEFAULT false NOT NULL -- Set default value to false
);

-- Create TABLE smartcontracts
CREATE TABLE smartcontracts (
    address VARCHAR(255) PRIMARY KEY,
    project_name VARCHAR(255),
    abi JSON,
    deployer_wallet VARCHAR(255),
    description TEXT,
    state_variables JSON -- Assuming map[string]string can be stored as JSON
);
