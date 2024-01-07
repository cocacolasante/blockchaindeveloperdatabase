package postgresrepo

import (
	"context"
	"database/sql"
	"log"
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

func (db *PostgresDb) AddWalletToDb(address string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	apikey := tools.GenerateApiKey()

	stmt := `insert into walletaccounts (wallet_address, credits_available, api_key ) values ($1, $2, $3)`

	_, err := db.Db.ExecContext(ctx, stmt, address, 0, apikey)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return apikey, nil

}

func(db *PostgresDb) GetWalletByAddress(address string) (*models.WalletAccount, error){
	log.Println("db call hit")
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	
	query := `SELECT wallet_address, credits_available 
			FROM walletaccounts
			WHERE wallet_address =  $1; `
	
	var wallet *models.WalletAccount

	err := db.Db.QueryRowContext(ctx, query, address).Scan(
		&wallet.WalletAddress,
		&wallet.CreditsAvailable,
	)
	if err != nil {
		return nil, err
	}

	return wallet, err

}
func(db *PostgresDb) AuthWalletByAddress(address string) (*models.WalletAccount, error){
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT (wallet_address, COALESCE(smart_contract_addresses, []), credits_available, api_key )
			FROM walletaccounts
			WHERE wallet_address = $1; `
	
	var wallet *models.WalletAccount

	err := db.Db.QueryRowContext(ctx, query, address).Scan(
		&wallet.WalletAddress,
		&wallet.SmartContractAddresses,
		&wallet.CreditsAvailable,
		&wallet.ApiKey,
	)

	if err != nil {
		return nil, err
	}

	return wallet, err
}
