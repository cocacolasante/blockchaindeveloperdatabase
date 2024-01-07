package models

type SmartContract struct {
	Address        string            `json:"address"`
	ProjectName    string            `json:"project_name"`
	Abi            ABI               `json:"abi"`
	DeployerWallet string            `json:"deployer_wallet"`
	Description    string            `json:"description"`
	StateVariables map[string]string `json:"state_variables"`
}

type ABI struct {
	Data any `json:"abi"`
}
