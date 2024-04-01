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

func TruncateTable() {
	_, err := datastore.EntClient.ExecContext(
		context.Background(),
		"DELETE FROM personal_accounts;"+
			"DELETE FROM sqlite_sequence WHERE name = 'personal_accounts';",
	)
	if err != nil {
		panic(err)
	}
}

func TestPersonalAccountRepository_FindAll(t *testing.T) {
	r := NewEntPersonalAccountRepository()

	all, err := r.FindAll(nil, nil)
	assert.Nil(t, err)
	assert.Empty(t, all)

	createPersonalAccountRandom(datastore.EntClient)

	all, err = r.FindAll(nil, nil)
	assert.Nil(t, err)
	assert.NotEmpty(t, all)

	TruncateTable()
}

func TestPersonalAccountRepository_FindById(t *testing.T) {
	r := NewEntPersonalAccountRepository()

	pa, err := r.FindById(1)
	assert.NotNil(t, err)
	assert.Empty(t, pa)

	random, err := createPersonalAccount(datastore.EntClient)
	assert.Nil(t, err)

	pa, err = r.FindById(1)
	assert.Nil(t, err)
	assert.NotEmpty(t, pa)

	assert.EqualValues(t, random.ID, pa.ID)

	TruncateTable()
}

func TestPersonalAccountRepository_Create(t *testing.T) {
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

	TruncateTable()
}
