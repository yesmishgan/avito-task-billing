package service

import (
	"cashbox"
	"cashbox/pkg/repository"
)

type Transaction interface {

}

type Info interface {
	GetBalance(username string) (cashbox.Account, error)
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