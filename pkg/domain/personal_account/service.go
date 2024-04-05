package domain

import (
	"Savings/ent"
	"Savings/pkg/repositories/filters"
	"Savings/utils"
)

type PersonalAccountService interface {
	FindAllPersonalAccounts(*utils.Paginator, *filters.PersonalAccountFilters) (ent.PersonalAccounts, error)
	FindPersonalAccountById(uint64) (*ent.PersonalAccount, error)
	CreatePersonalAccount(*ent.PersonalAccount) (*ent.PersonalAccount, error)
}

type personalAccountService struct {
	repo PersonalAccountRepository
}

func NewPersonalAccountService(repository PersonalAccountRepository) PersonalAccountService {
	return &personalAccountService{repo: repository}
}

func (p *personalAccountService) FindAllPersonalAccounts(paginator *utils.Paginator, filters *filters.PersonalAccountFilters) (ent.PersonalAccounts, error) {
	return p.repo.FindAll(paginator, filters)
}

func (p *personalAccountService) FindPersonalAccountById(id uint64) (*ent.PersonalAccount, error) {
	return p.repo.FindById(id)
}

func (p *personalAccountService) CreatePersonalAccount(data *ent.PersonalAccount) (*ent.PersonalAccount, error) {
	return p.repo.Create(data)
}

func (p *personalAccountService) FindPersonalAccountsByAccountId(id uint64) (ent.PersonalAccounts, error) {
	return p.repo.FindByAccountId(id)
}
