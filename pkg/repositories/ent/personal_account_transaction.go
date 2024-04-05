package entRepo

import (
	"Savings/ent"
	"Savings/ent/personalaccounttransaction"
	"Savings/pkg/datastore"
	domain "Savings/pkg/domain/personal_account_transaction"
	"Savings/pkg/repositories/filters"
	"Savings/utils"
	"context"
	"github.com/spf13/viper"
)

type personalAccountTransactionRepository struct {
	client *ent.PersonalAccountTransactionClient
}

func NewEntPersonalAccountTransactionRepository() domain.PersonalAccountTransactionRepository {
	c := datastore.EntClient
	if viper.GetBool("APP_DEBUG") {
		c = c.Debug()
	}
	return &personalAccountTransactionRepository{
		client: c.PersonalAccountTransaction,
	}
}

func (p *personalAccountTransactionRepository) FindAll(paginator *utils.Paginator, filters *filters.PersonalAccountTransactionFilters) (transactions ent.PersonalAccountTransactions, err error) {
	q := p.client.Query()

	if filters != nil {
		if filters.PersonalAccountId != 0 {
			q.Where(personalaccounttransaction.PersonalAccountID(filters.PersonalAccountId))
		}
		if filters.Type != "" {
			q.Where(personalaccounttransaction.Type(filters.Type))
		}
		if filters.Status != "" {
			q.Where(personalaccounttransaction.Status(filters.Status))
		}
	}

	return q.Limit(paginator.PageSize()).Offset(paginator.Offset()).All(context.Background())
}

func (p *personalAccountTransactionRepository) FindById(id uint64) (transaction *ent.PersonalAccountTransaction, err error) {
	return p.client.Query().Where(personalaccounttransaction.ID(id)).First(context.Background())
}
