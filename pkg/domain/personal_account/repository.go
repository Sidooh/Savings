package domain

import (
	"Savings/ent"
	"Savings/pkg/repositories/filters"
	"Savings/utils"
)

type PersonalAccountRepository interface {
	FindAll(*utils.Paginator, *filters.PersonalAccountFilters) (ent.PersonalAccounts, error)
	FindById(uint64) (*ent.PersonalAccount, error)
	Create(*ent.PersonalAccount) (*ent.PersonalAccount, error)

	Credit(id uint64, amount float32, description string) error
	Debit(id uint64, amount float32, description string) error

	CalculateInterest() error
	AllocateInterest() error
}
