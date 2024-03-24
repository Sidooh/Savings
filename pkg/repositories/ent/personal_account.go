package entRepo

import (
	"Savings/ent"
	"Savings/pkg/datastore"
	domain "Savings/pkg/domain/personal_account"
	"context"
)

type personalAccountRepository struct {
	client *ent.PersonalAccountClient
}

func NewEntPersonalAccountRepository() domain.PersonalAccountRepository {
	return &personalAccountRepository{client: datastore.EntClient.PersonalAccount}
}

func (p personalAccountRepository) FindAll() (accounts []*ent.PersonalAccount, err error) {
	accounts, err = p.client.Query().All(context.Background())
	return
}
