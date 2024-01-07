package models

type SmartContract struct {
	Address        string            `json:"address"`
	ProjectName    string            `json:"project_name"`
	Abi            any               `json:"abi,omitempty"`
	DeployerWallet string            `json:"deployer_wallet"`
	Description    string            `json:"description"`
	StateVariables map[string]string `json:"state_variables,omitempty"`
}

