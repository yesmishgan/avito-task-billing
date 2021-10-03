package service

import (
	"cashbox"
	"cashbox/pkg/repository"
)

type Transaction interface {
	Write(bill cashbox.Bill) error
	Transfer(transfer cashbox.Transfer) error
}

type Info interface {
	GetBalance(user cashbox.User) (cashbox.Account, error)
	//GetActions(username string)
}

type Service struct {
	Transaction
	Info
}

func NewService(repos *repository.Repository) *Service{
	return &Service{
		Info: NewInfoService(repos.Info),
		Transaction: NewTransactionService(repos.Transaction),
	}
}