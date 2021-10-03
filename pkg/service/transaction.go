package service

import (
	"cashbox"
	"cashbox/pkg/repository"
)

type TransactionService struct {
	repo repository.Transaction
}

func NewTransactionService(repo repository.Transaction) *TransactionService {
	return &TransactionService{repo: repo}
}

func (t *TransactionService) Write(bill cashbox.Bill) error{
	return t.repo.Write(bill)
}

func (t *TransactionService) Transfer(transfer cashbox.Transfer) error{
	return t.repo.Transfer(transfer)
}
