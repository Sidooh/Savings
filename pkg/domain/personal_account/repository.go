package domain

import "Savings/ent"

type PersonalAccountRepository interface {
	FindAll() (accounts []*ent.PersonalAccount, err error)
	FindById(id uint64) (account *ent.PersonalAccount, err error)
	Create(data *ent.PersonalAccount) (account *ent.PersonalAccount, err error)
}
