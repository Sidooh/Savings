package domain

import (
	"Savings/ent"
)

type PersonalAccountService interface {
	FindAllPersonalAccounts() ([]*ent.PersonalAccount, error)
}

type personalAccountService struct {
	repo PersonalAccountRepository
}

func NewPersonalAccountService(repository PersonalAccountRepository) PersonalAccountService {
	return &personalAccountService{repo: repository}
}

func (p *personalAccountService) FindAllPersonalAccounts() ([]*ent.PersonalAccount, error) {
	return p.repo.FindAll()
}
