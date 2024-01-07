package models

type WalletAccount struct {
	WalletAddress          string   `json:"wallet_address"`
	SmartContractAddresses []string `json:"smart_contract_addresses"`
	CreditsAvailable       int      `json:"credits_available"`
	ApiKey                 string   `json:"api_key"`
}
