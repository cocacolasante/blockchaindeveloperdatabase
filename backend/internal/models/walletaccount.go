package models

import "math/big"

// UPDATE CREDITS AVAIL TYPING TO EITHER BIG INT OR INT AND UPDATE DATABASE
type WalletAccount struct {
	WalletAddress          string   `json:"wallet_address"`
	Email                  string   `json:"email"`
	Password               string   `json:"password"`
	SmartContractAddresses []string `json:"smart_contract_addresses"`
	CreditsAvailable       *big.Int `json:"credits_available"`
	ApiKey                 string   `json:"api_key"`
}
