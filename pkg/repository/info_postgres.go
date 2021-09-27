package repository

import (
	"cashbox"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type InfoPostgres struct {
	db *sqlx.DB
}

func NewInfoPostgres(db *sqlx.DB) *InfoPostgres {
	return &InfoPostgres{db: db}
}

func (i *InfoPostgres) GetBalance(username string, currency string) (cashbox.Account, error) {
	var account cashbox.Account
	fmt.Println("LET'S GO")
	query := fmt.Sprintf("SELECT balance FROM %s WHERE username = $1", clientsTable)
	err := i.db.Get(&account, query, username)

	return account, err
}
