package domain

import (
	"Savings/ent"
)

type PersonalAccountService interface {
	FindAllPersonalAccounts() ([]*ent.PersonalAccount, error)
	FindPersonalAccountById(id uint64) (*ent.PersonalAccount, error)
	CreateUser(user *ent.PersonalAccount) (*ent.PersonalAccount, error)
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

func (p *personalAccountService) FindPersonalAccountById(id uint64) (*ent.PersonalAccount, error) {
	return p.repo.FindById(id)
}

func (p *personalAccountService) CreateUser(user *ent.PersonalAccount) (*ent.PersonalAccount, error) {
	return p.repo.Create(user)
}
