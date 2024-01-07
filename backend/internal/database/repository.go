package repository

import (
	"database/sql"

	"github.com/cocacolasante/blockchaindeveloperdatabase/internal/models"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	AddWalletToDb(address string) (*models.WalletAccount, error)
}
