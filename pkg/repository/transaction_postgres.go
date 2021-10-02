package repository

import (
	"cashbox"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type TransactionPostgres struct {
	db *sqlx.DB
}

func NewTransactionPostgres(db *sqlx.DB) *TransactionPostgres {
	return &TransactionPostgres{db: db}
}

func (t *TransactionPostgres) Write(bill cashbox.Bill) error{

	if bill.Amount <= 0 {
		return errors.New("amount must be positive")
	}

	tx, err := t.db.Beginx()
	if err != nil{
		return err
	}

	var id, idT int
	var query string

	if bill.Flag {
		query = fmt.Sprintf(`SELECT write($1, $2);`)
		row := tx.QueryRow(query, bill.Username, bill.Amount)
		if err = row.Scan(&id); err != nil {
			tx.Rollback()
			return err
		}
	} else {
		query = fmt.Sprintf("SELECT id FROM %s WHERE username = $1", clientsTable)
		if err := t.db.Get(&id, query, bill.Username); err != nil{
			tx.Rollback()
			return err
		}

		query = fmt.Sprintf(`SELECT id FROM %s WHERE id = $1 AND balance >= $2;`, clientsTable)
		if err := t.db.Get(&id, query, id, bill.Amount); err != nil{
			tx.Rollback()
			return errors.New("write-off is greater than balance")
		}

		query = fmt.Sprintf(` UPDATE %s
										SET balance = balance - $2
        								WHERE id = $1;`,
										clientsTable)
		_, err := tx.Exec(query, id, bill.Amount)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	query = fmt.Sprintf(`INSERT INTO %s (description, sender, destination, amount)
									VALUES ($1, $2, $3, $4)
									RETURNING id;`, transactionsTable)

	var row *sql.Row

	if bill.Flag{
		row = tx.QueryRow(query, bill.Description, 0, id, bill.Amount)
	} else {
		row = tx.QueryRow(query, bill.Description, id, 0, bill.Amount)
	}

	if err = row.Scan(&idT); err != nil {
		tx.Rollback()
		return err
	}

	query = fmt.Sprintf(`INSERT INTO %s (client_id, transaction_id)
							VALUES ($1, $2)`, clientTransactionsTable)

	_, err = tx.Exec(query, id, idT)
	if err != nil{
		tx.Rollback()
		return err
	}

	return tx.Commit()
}