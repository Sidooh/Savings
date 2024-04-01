package domain

import "Savings/ent"

type PersonalAccountTransactionRepository interface {
	FindAll() ([]*ent.PersonalAccountTransaction, error)
	FindById(uint64) (*ent.PersonalAccountTransaction, error)
	Create(*ent.PersonalAccountTransaction) (*ent.PersonalAccountTransaction, error)
}
