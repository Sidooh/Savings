package domain

import (
	"Savings/ent"
	"Savings/pkg/repositories/filters"
	"Savings/utils"
)

type PersonalAccountTransactionRepository interface {
	FindAll(*utils.Paginator, *filters.PersonalAccountTransactionFilters) (ent.PersonalAccountTransactions, error)
	FindById(uint64) (*ent.PersonalAccountTransaction, error)
}
