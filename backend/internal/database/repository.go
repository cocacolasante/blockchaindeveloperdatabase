package repository

import (
	"database/sql"

	"github.com/cocacolasante/blockchaindeveloperdatabase/internal/models"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	AddWalletToDb(address, email, password string) (string, error)
	GetWalletByAddress(address string) (*models.WalletAccount, error)
	AdminGetWalletAccount(address string) (*models.WalletAccount, error)
	UpdateAPIKey(address, key string) (string, error)
	AddSmartContractToAccountDb(contract models.SmartContract, id string) (error)
	GetSmartContract(address string) (*models.SmartContract, error)
	UpdateSmartContractToAccountDb(updateAddress string, contract models.SmartContract) (error)
	DeleteSmartContract(address, userAddress string)(error)
	GetAllSmartContractInWalletAccounts(userAddress string) (*[]string, error)
	GetAllFullScInWallet(userAddress string) (*[]models.SmartContract, error)
	AdminGetWalletAccountByEmail(email string) (*models.WalletAccount, error)
	ActivateAccount(walletAddress string) (error)
}
