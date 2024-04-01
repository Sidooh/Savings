package domain

import (
	"Savings/ent"
	"Savings/pkg/repositories"
	"Savings/utils"
)

type PersonalAccountRepository interface {
	FindAll(*utils.Paginator, *repositories.PersonalAccountFilters) (ent.PersonalAccounts, error)
	FindById(uint64) (*ent.PersonalAccount, error)
	Create(*ent.PersonalAccount) (*ent.PersonalAccount, error)

	FindByAccountId(uint64) (ent.PersonalAccounts, error)
}
