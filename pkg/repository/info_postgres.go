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

func (i *InfoPostgres) GetBalance(user cashbox.User) (cashbox.Account, error) {
	var account cashbox.Account

	query := fmt.Sprintf("SELECT balance FROM %s WHERE username = $1", clientsTable)
	err := i.db.Get(&account, query, user.Username)

	return account, err
}
