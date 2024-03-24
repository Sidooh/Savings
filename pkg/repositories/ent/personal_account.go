package entRepo

import (
	"Savings/ent"
	"Savings/ent/personalaccount"
	"Savings/pkg/datastore"
	domain "Savings/pkg/domain/personal_account"
	"context"
)

type personalAccountRepository struct {
	client *ent.PersonalAccountClient
	query  *ent.PersonalAccountQuery
}

func NewEntPersonalAccountRepository() domain.PersonalAccountRepository {
	return &personalAccountRepository{
		client: datastore.EntClient.PersonalAccount,
		query:  datastore.EntClient.PersonalAccount.Query(),
	}
}

func (p personalAccountRepository) FindAll() (accounts []*ent.PersonalAccount, err error) {
	return p.client.Query().All(context.Background())
}

func (p personalAccountRepository) FindById(id uint64) (account *ent.PersonalAccount, err error) {
	return p.query.Where(personalaccount.ID(id)).First(context.Background())
}

func (p personalAccountRepository) Create(data *ent.PersonalAccount) (account *ent.PersonalAccount, err error) {
	return p.client.Create().SetAccountID(data.AccountID).SetType(data.Type).Save(context.Background())
}
