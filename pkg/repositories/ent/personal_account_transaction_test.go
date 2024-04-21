package entRepo

import (
	"Savings/ent"
	"Savings/pkg/datastore"
	"context"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	viper.Set("APP_ENV", "TEST")
	viper.Set("DB_DSN", "file:ent?mode=memory&cache=shared&_fk=1")
	datastore.Init()
}

func createPersonalAccountTransaction(c *ent.Client) (*ent.PersonalAccountTransaction, error) {
	acc, _ := createPersonalAccount(c)
	return c.PersonalAccountTransaction.Create().
		SetAccount(acc).
		SetType("CREDIT").
		SetStatus("COMPLETED").
		Save(context.Background())
}

func TestPersonalAccountTransactionRepository_FindAll(t *testing.T) {
	r := NewEntPersonalAccountTransactionRepository()

	all, err := r.FindAll(nil, nil)
	assert.Nil(t, err)
	assert.Empty(t, all)

	createPersonalAccountTransaction(datastore.EntClient)

	all, err = r.FindAll(nil, nil)
	assert.Nil(t, err)
	assert.NotEmpty(t, all)

	TruncateTable("personal_account_transactions")
}

func TestPersonalAccountTransactionRepository_FindById(t *testing.T) {
	r := NewEntPersonalAccountTransactionRepository()

	pa, err := r.FindById(1)
	assert.NotNil(t, err)
	assert.Empty(t, pa)

	random, err := createPersonalAccountTransaction(datastore.EntClient)
	assert.Nil(t, err)

	pa, err = r.FindById(1)
	assert.Nil(t, err)
	assert.NotEmpty(t, pa)

	assert.EqualValues(t, random.ID, pa.ID)

	TruncateTable("personal_account_transactions")
}

func TestPersonalAccountRepository_Cred(t *testing.T) {
	r := NewEntPersonalAccountRepository()

	acc := ent.PersonalAccount{
		AccountID: 1,
		Type:      "",
	}

	pa, err := r.FindById(1)
	assert.NotNil(t, err)
	assert.Empty(t, pa)

	created, err := r.Create(&acc)
	assert.Nil(t, err)
	assert.EqualValues(t, 1, created.ID)
	assert.EqualValues(t, 1, created.AccountID)
	assert.EqualValues(t, "", created.Type)

	created, err = r.Create(&acc)
	assert.Nil(t, created)
	assert.NotNil(t, err)

	TruncateTable("personal_account_transactions")
}
