package postgresrepo

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"strings"
	"time"

	"github.com/cocacolasante/blockchaindeveloperdatabase/internal/models"
	"github.com/cocacolasante/blockchaindeveloperdatabase/internal/tools"
)

type PostgresDb struct {
	Db *sql.DB
}

const dbTimeout = time.Second * 3

func (m *PostgresDb) Connection() *sql.DB {
	return m.Db
}

func (db *PostgresDb) AddWalletToDb(wallet *models.WalletAccount) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	// @todo use lower case to add to db
	// lowerWall := strings.ToLower(wallet.WalletAddress)
	// lowerEmail := strings.ToLower(wallet.Email)

	apikey := tools.GenerateApiKey()

	stmt := `insert into walletaccounts (wallet_address, api_key, email, password, credits_available ) values ($1, $2, $3, $4, $5)`

	_, err := db.Db.ExecContext(ctx, stmt, strings.ToLower(wallet.WalletAddress), apikey, strings.ToLower(wallet.Email), wallet.Password, wallet.CreditsAvailable)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return apikey, nil

}

func (db *PostgresDb) GetWalletByAddress(address string) (*models.WalletAccount, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// @todo make wallet address lowercase before searching
	// lowerWall := strings.ToLower(address)

	query := `SELECT wallet_address, credits_available, email 
			FROM walletaccounts
			WHERE wallet_address =  $1; `

	var wallet models.WalletAccount

	err := db.Db.QueryRowContext(ctx, query, strings.ToLower(address)).Scan(
		&wallet.WalletAddress,
		&wallet.CreditsAvailable,
		&wallet.Email,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return &models.WalletAccount{}, nil
	}
	if err != nil {
		return nil, err
	}

	return &wallet, err

}

func (db *PostgresDb) AdminGetWalletAccount(address string) (*models.WalletAccount, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	
	// @todo make wallet address lowercase before searching
	// lowerWall := strings.ToLower(address)


	query := ` SELECT
            wallet_address,
            credits_available,
			email,
            api_key,
			activated,
            COALESCE(smart_contract_addresses, '{}'::VARCHAR[]) AS smart_contract_addresses
        FROM walletaccounts
        WHERE wallet_address = $1; `

	var wallet models.WalletAccount
	var smartContractsStr string
	err := db.Db.QueryRowContext(ctx, query, strings.ToLower(address)).Scan(
		&wallet.WalletAddress,
		&wallet.CreditsAvailable,
		&wallet.Email,
		&wallet.ApiKey,
		&wallet.Active,
		&smartContractsStr,
	)

	if err != nil {
		return nil, err
	}

	wallet.SmartContractAddresses = append(wallet.SmartContractAddresses, smartContractsStr)

	return &wallet, err
}
func (db *PostgresDb) AdminGetWalletAccountByEmail(email string) (*models.WalletAccount, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

		// @todo make wallet address lowercase before searching
	// loweremail := strings.ToLower(email)


	query := `SELECT
            wallet_address,
            credits_available,
			email,
			password,
            api_key,
			activated,
            COALESCE(smart_contract_addresses, '{}'::VARCHAR[]) AS smart_contract_addresses
        FROM walletaccounts
        WHERE email = $1; `

	var wallet models.WalletAccount
	var smartContractsStr string
	err := db.Db.QueryRowContext(ctx, query, strings.ToLower(email)).Scan(
		&wallet.WalletAddress,
		&wallet.CreditsAvailable,
		&wallet.Email,
		&wallet.Password,
		&wallet.ApiKey,
		&wallet.Active,
		&smartContractsStr,
	)

	if err != nil {
		return nil, err
	}
	
		
	wallet.SmartContractAddresses = append(wallet.SmartContractAddresses, smartContractsStr)
	

	return &wallet, err
}

func (db *PostgresDb) UpdateAPIKey(address, key string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `UPDATE walletaccounts SET api_key = $1 WHERE wallet_address = $2 returning api_key`

	var newKey string
	err := db.Db.QueryRowContext(ctx, stmt, key, address).Scan(&newKey)
	if err != nil {
		return "", err
	}
	return newKey, nil

}

func (db *PostgresDb) GetSmartContract(address string) (*models.SmartContract, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	var contract models.SmartContract
	var statevar any
	query := `SELECT address, project_name, abi, deployer_wallet, description, state_variables FROM smartcontracts WHERE address = $1`

	err := db.Db.QueryRowContext(ctx, query, strings.ToLower(address)).Scan(
		&contract.Address,
		&contract.ProjectName,
		&contract.Abi,
		&contract.DeployerWallet,
		&contract.Description,
		&statevar,
	)
	// TO DO -- MAP STATEVAR TO THE STATEVARIABLES MAP[STRING]STRING

	if err != nil {
		return nil, err
	}

	return &contract, nil

}

func (db *PostgresDb) GetAllSmartContractInWalletAccounts(userAddress string) (*[]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var contracts []string

	firstQuery := `SELECT smart_contract_addresses FROM walletaccounts WHERE wallet_address = $1`
	rows, err := db.Db.QueryContext(ctx, firstQuery, strings.ToLower(userAddress))
	if err != nil {
		log.Println("query", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var smartContractsStr string
		if err := rows.Scan(&smartContractsStr); err != nil {
			log.Println("scan", err)
			return nil, err
		}
		addresses := strings.Split(strings.Trim(smartContractsStr, "{}"), ",")

		// Trim spaces from each address
		for i, address := range addresses {
			addresses[i] = strings.TrimSpace(address)
		}

		contracts = append(contracts, addresses...)
	}

	log.Println("final contracts:", contracts)

	return &contracts, nil
}

func (db *PostgresDb) GetAllFullScInWallet(userAddress string) (*[]models.SmartContract, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	addresses, err := db.GetAllSmartContractInWalletAccounts(userAddress)
	if err != nil {
		log.Println("query", err)
		return nil, err
	}
	quotedAddresses := make([]string, len(*addresses))
    for i, address := range *addresses {
        quotedAddresses[i] = "'" + strings.ToLower(address) + "'"
    }
	query := `SELECT address, project_name, abi, deployer_wallet, description FROM smartcontracts WHERE address IN (` + strings.Join((quotedAddresses), ",") + `);`
	rows, err := db.Db.QueryContext(ctx, query)
	if err != nil {
		log.Println("query", err)
		return nil, err
	}
	defer rows.Close()
	var smartContracts []models.SmartContract

	// Iterate over the rows and scan each result into a SmartContract struct
	for rows.Next() {
		var sc models.SmartContract
		if err := rows.Scan(&sc.Address, &sc.ProjectName, &sc.Abi, &sc.DeployerWallet, &sc.Description); err != nil {
			log.Println("scan", err)
			return nil, err
		}
		smartContracts = append(smartContracts, sc)
	}

	log.Println("final smart contracts:", smartContracts)

	return &smartContracts, nil
}

func (db *PostgresDb) AddSmartContractToAccountDb(contract models.SmartContract, id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	// @todo make all wallet addresses lower case prior to adding to db, - case sensitive is affecting exequery
	stmt := `INSERT INTO smartcontracts (address, project_name, deployer_wallet, description) values($1, $2, $3, $4);`

	_, err := db.Db.ExecContext(ctx, stmt, strings.ToLower(contract.Address), contract.ProjectName, strings.ToLower(contract.DeployerWallet), contract.Description)
	if err != nil {
		return errors.New("error adding into smartcontracts: " + err.Error())
	}
	
	stmt = `UPDATE walletaccounts
		SET smart_contract_addresses = smart_contract_addresses || ARRAY[$1]
		WHERE wallet_address = $2;`

	_, err = db.Db.ExecContext(ctx, stmt, strings.ToLower(contract.Address), strings.ToLower(id))
	if err != nil {
		return errors.New("error adding into walletaccounts: " + err.Error())
	}
	return nil
}

func (db *PostgresDb) UpdateSmartContractToAccountDb(updateAddress string, contract models.SmartContract) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
        UPDATE smartcontracts
        SET
            project_name = $2,
            abi = $3,
            description = $4,
            state_variables = $5
        WHERE address = $1;
    `

	_, err := db.Db.ExecContext(ctx, query, strings.ToLower(updateAddress), contract.ProjectName, contract.Abi, contract.Description, contract.StateVariables)
	if err != nil {
		return err
	}

	return nil
}

func (db *PostgresDb) DeleteSmartContract(address, userAddress string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `delete from smartcontracts where address =$1;`
	_, err := db.Db.ExecContext(ctx, stmt, strings.ToLower(address))
	if err != nil {
		return err
	}
	stmt = `UPDATE walletaccounts
		SET smart_contract_addresses = array_remove(smart_contract_addresses, $1)
		WHERE wallet_address = $2;
	`
	_, err = db.Db.ExecContext(ctx, stmt, address, strings.ToLower(userAddress))
	if err != nil {
		return err
	}

	return nil
}


func(db *PostgresDb) ActivateAccount(walletAddress string) (error){
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	stmt := `UPDATE walletaccounts
			SET activated = $1
			WHERE wallet_address = $2;`
	_, err := db.Db.ExecContext(ctx, stmt, true, strings.ToLower(walletAddress))
	if err != nil {
		return err
	}

	return nil
}

