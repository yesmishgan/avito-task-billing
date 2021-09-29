package repository

import (
	"cashbox"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type TransactionPostgres struct {
	db *sqlx.DB
}

func NewTransactionPostgres(db *sqlx.DB) *TransactionPostgres {
	return &TransactionPostgres{db: db}
}

func (t *TransactionPostgres) Write(bill cashbox.Bill) error{

	if !bill.Flag {
		panic("make!")
	}

	tx, err := t.db.Beginx()
	if err != nil{
		return err
	}

	var id int
	query := fmt.Sprintf(`SELECT write($1, $2);`)
	row := tx.QueryRow(query, bill.Username, bill.Amount)
	if err = row.Scan(&id); err != nil {
		tx.Rollback()
		return err
	}

	query = fmt.Sprintf(`INSERT INTO %s (description, sender, destination, amount)
									VALUES ($1, $2, $3, $4)
									RETURNING id;`, transactionsTable)
	row = tx.QueryRow(query, bill.Description, id, id, bill.Amount)
	var idT int
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