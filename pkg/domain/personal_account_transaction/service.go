package domain

import (
	"Savings/ent"
)

type PersonalAccountTransactionService interface {
	FindAllPersonalAccountTransactions() (ent.PersonalAccountTransactions, error)
	FindPersonalAccountTransactionById(uint64) (*ent.PersonalAccountTransaction, error)
	CreatePersonalAccountTransaction(*ent.PersonalAccountTransaction) (*ent.PersonalAccountTransaction, error)
}

type personalAccountTransactionService struct {
	repo PersonalAccountTransactionRepository
}

func NewPersonalAccountTransactionService(repository PersonalAccountTransactionRepository) PersonalAccountTransactionService {
	return &personalAccountTransactionService{repo: repository}
}

func (p *personalAccountTransactionService) FindAllPersonalAccountTransactions() (ent.PersonalAccountTransactions, error) {
	return p.repo.FindAll()
}

func (p *personalAccountTransactionService) FindPersonalAccountTransactionById(id uint64) (*ent.PersonalAccountTransaction, error) {
	return p.repo.FindById(id)
}

func (p *personalAccountTransactionService) CreatePersonalAccountTransaction(data *ent.PersonalAccountTransaction) (*ent.PersonalAccountTransaction, error) {
	return p.repo.Create(data)
}
