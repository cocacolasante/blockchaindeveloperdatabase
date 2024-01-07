package models

type WalletAccount struct {
	WalletAddress          string   `json:"wallet_address"`
	SmartContractAddresses []string `json:"smart_contract_addresses,omitempty"`
	CreditsAvailable       int      `json:"credits_available,omitempty"`
	ApiKey                 string   `json:"api_key,omitempty"`
}
