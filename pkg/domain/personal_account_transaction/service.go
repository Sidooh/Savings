package domain

import (
	"Savings/ent"
	"Savings/pkg/repositories/filters"
	"Savings/utils"
)

type PersonalAccountTransactionService interface {
	FindAllPersonalAccountTransactions(paginator *utils.Paginator, filters *filters.PersonalAccountTransactionFilters) (ent.PersonalAccountTransactions, error)
	FindPersonalAccountTransactionById(uint64) (*ent.PersonalAccountTransaction, error)
	//CreatePersonalAccountTransaction(*ent.PersonalAccountTransaction) (*ent.PersonalAccountTransaction, error)
}

type personalAccountTransactionService struct {
	repo PersonalAccountTransactionRepository
}

func NewPersonalAccountTransactionService(repository PersonalAccountTransactionRepository) PersonalAccountTransactionService {
	return &personalAccountTransactionService{repo: repository}
}

func (p *personalAccountTransactionService) FindAllPersonalAccountTransactions(paginator *utils.Paginator, filters *filters.PersonalAccountTransactionFilters) (ent.PersonalAccountTransactions, error) {
	return p.repo.FindAll(paginator, filters)
}

func (p *personalAccountTransactionService) FindPersonalAccountTransactionById(id uint64) (*ent.PersonalAccountTransaction, error) {
	return p.repo.FindById(id)
}
