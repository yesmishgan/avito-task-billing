package repository

import (
	"cashbox"
	"github.com/jmoiron/sqlx"
)

type Transaction interface {
	Write(bill cashbox.Bill) error
}

type Info interface {
	GetBalance(user cashbox.User) (cashbox.Account, error)
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
