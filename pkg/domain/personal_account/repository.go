package domain

import "Savings/ent"

type PersonalAccountRepository interface {
	FindAll() ([]*ent.PersonalAccount, error)

	//Create(user *ent.PersonalAccount) error
	//FindById(id string) (*User, error)

	//Save(user *User) error
}
