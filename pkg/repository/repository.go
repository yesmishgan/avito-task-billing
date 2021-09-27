package repository

import (
	"cashbox"
	"github.com/jmoiron/sqlx"
)

type Transaction interface {

}

type Info interface {
	GetBalance(username string, currency string) (cashbox.Account, error)
}

type Repository struct {
	Transaction
	Info
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Transaction: NewTransactionPostgres(db),
		Info: NewInfoPostgres(db),
	}
}
