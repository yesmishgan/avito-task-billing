package service

import (
	"cashbox"
	"cashbox/pkg/repository"
)

type InfoService struct {
	repo repository.Info
}

func NewInfoService(repo repository.Info) *InfoService {
	return &InfoService{repo: repo}
}

func (i *InfoService) GetBalance(username string, currency string) (cashbox.Account, error) {
	return i.repo.GetBalance(username, currency)
}