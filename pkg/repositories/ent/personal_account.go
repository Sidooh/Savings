package entRepo

import (
	"Savings/ent"
	"Savings/ent/personalaccount"
	"Savings/pkg/datastore"
	domain "Savings/pkg/domain/personal_account"
	"Savings/pkg/repositories"
	"Savings/utils"
	"context"
	"github.com/spf13/viper"
)

type personalAccountRepository struct {
	client *ent.PersonalAccountClient
}

func NewEntPersonalAccountRepository() domain.PersonalAccountRepository {
	c := datastore.EntClient
	if viper.GetBool("APP_DEBUG") {
		c = c.Debug()
	}
	return &personalAccountRepository{
		client: c.PersonalAccount,
	}
}

func (p personalAccountRepository) FindAll(paginator *utils.Paginator, filters *repositories.PersonalAccountFilters) (accounts ent.PersonalAccounts, err error) {
	q := p.client.Query()

	if filters != nil && filters.AccountId != 0 {
		q = q.Where(personalaccount.AccountID(filters.AccountId))
	}

	return q.Limit(paginator.PageSize()).Offset(paginator.Offset()).All(context.Background())
}

func (p personalAccountRepository) FindById(id uint64) (account *ent.PersonalAccount, err error) {
	return p.client.Query().Where(personalaccount.ID(id)).First(context.Background())
}

func (p personalAccountRepository) Create(data *ent.PersonalAccount) (account *ent.PersonalAccount, err error) {
	return p.client.Create().SetAccountID(data.AccountID).SetType(data.Type).Save(context.Background())
}

func (p personalAccountRepository) FindByAccountId(id uint64) (accounts ent.PersonalAccounts, err error) {
	return p.client.Query().Where(personalaccount.AccountID(id)).All(context.Background())
}
