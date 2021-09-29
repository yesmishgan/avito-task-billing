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

func (i *InfoService) GetBalance(user cashbox.User) (cashbox.Account, error) {
	return i.repo.GetBalance(user)
}