package postgresrepo

import (
	"context"
	"database/sql"
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

func (db *PostgresDb) AddWalletToDb(address string) (*models.WalletAccount, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	var wallet *models.WalletAccount

	apikey := tools.GenerateApiKey()

	wallet.ApiKey = apikey
	wallet.CreditsAvailable = 0
	wallet.WalletAddress = address

	stmt := `insert into walletaccounts (walletaddress, creditsavailable, apikey ) values ($1, $2, $3)`

	_, err := db.Db.ExecContext(ctx, stmt, wallet.WalletAddress, wallet.CreditsAvailable, wallet.ApiKey)
	if err != nil {
		return nil, err
	}

	return wallet, nil

}


