package entRepo

import (
	"Savings/ent"
	"Savings/pkg/datastore"
	"Savings/utils"
	internal_errors "Savings/utils/errors"
	"context"
	"fmt"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func init() {
	viper.Set("APP_ENV", "TEST")
	viper.Set("DB_DSN", "file:ent?mode=memory&cache=shared&_fk=1")
	datastore.Init()
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

	TruncateTable("personal_accounts")
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

	TruncateTable("personal_accounts")
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

	TruncateTable("personal_accounts")
}

func TestPersonalAccountRepository_Credit(t *testing.T) {
	r := NewEntPersonalAccountRepository()
	r2 := NewEntPersonalAccountTransactionRepository()

	account, err := createPersonalAccount(datastore.EntClient)
	assert.Nil(t, err)

	err = r.Credit(account.ID, 10, "Test Credit")
	assert.Nil(t, err)

	acc, err := r.FindById(account.ID)
	assert.Nil(t, err)
	assert.NotNil(t, acc)
	assert.EqualValues(t, account.ID, acc.ID)
	assert.NotEqualValues(t, account.Balance, acc.Balance)
	assert.EqualValues(t, 10, acc.Balance)

	tx, err := r2.FindAll(nil, nil)
	assert.Nil(t, err)
	assert.EqualValues(t, 1, len(tx))
	assert.EqualValues(t, acc.ID, tx[0].PersonalAccountID)
	assert.EqualValues(t, acc.Balance, tx[0].Balance)

	TruncateTable("personal_account_transactions")
	TruncateTable("personal_accounts")
}

func TestPersonalAccountRepository_Debit(t *testing.T) {
	r := NewEntPersonalAccountRepository()
	r2 := NewEntPersonalAccountTransactionRepository()

	account, err := createPersonalAccount(datastore.EntClient)
	assert.Nil(t, err)

	// Throw error on negative debit
	err = r.Debit(account.ID, 10, "Test Debit")
	assert.NotNil(t, err)

	account, err = account.Update().AddBalance(10).Save(context.Background())
	assert.Nil(t, err)

	err = r.Debit(account.ID, 10, "Test Debit")
	assert.Nil(t, err)

	acc, err := r.FindById(account.ID)
	assert.Nil(t, err)
	assert.NotNil(t, acc)
	assert.EqualValues(t, account.ID, acc.ID)
	assert.NotEqualValues(t, account.Balance, acc.Balance)
	assert.EqualValues(t, 0, acc.Balance)

	tx, err := r2.FindAll(nil, nil)
	assert.Nil(t, err)
	assert.EqualValues(t, 1, len(tx))
	assert.EqualValues(t, acc.ID, tx[0].PersonalAccountID)
	assert.EqualValues(t, acc.Balance, tx[0].Balance)

	TruncateTable("personal_account_transactions")
	TruncateTable("personal_accounts")
}

func TestPersonalAccountRepository_CalculateInterest(t *testing.T) {
	viper.Set("APR", 9)
	r := NewEntPersonalAccountRepository()
	jr := NewEntJobRepository()

	// TEST successful computation
	jobs, err := jr.FindAll(nil, nil)
	assert.Len(t, jobs, 0)

	account, err := createPersonalAccount(datastore.EntClient)
	assert.Nil(t, err)

	err = r.CalculateInterest()
	assert.Nil(t, err)

	jobs, err = jr.FindAll(nil, nil)
	assert.Nil(t, err)
	assert.Len(t, jobs, 1)
	assert.EqualValues(t, 1, jobs[0].TotalProcessed)
	assert.EqualValues(t, "COMPLETED", jobs[0].Status)

	// TEST does not recalculate if completed
	err = r.CalculateInterest()
	assert.Error(t, err, internal_errors.JobAlreadyCompleted)

	jobs2, err := jr.FindAll(nil, nil)
	assert.Nil(t, err)
	assert.Len(t, jobs, 1)
	assert.EqualValues(t, 1, jobs[0].TotalProcessed)
	assert.EqualValues(t, "COMPLETED", jobs[0].Status)
	assert.EqualValues(t, jobs[0].UpdatedAt, jobs2[0].UpdatedAt)

	// TEST Interest calculation
	TruncateTable("jobs")
	viper.Set("APR", 12)

	acc, err := r.FindById(account.AccountID)
	assert.Nil(t, err)
	assert.EqualValues(t, 0, acc.Balance)
	assert.EqualValues(t, 0, acc.Interest)

	acc, err = acc.Update().AddBalance(100).Save(context.Background())
	assert.Nil(t, err)

	err = r.CalculateInterest()
	assert.Nil(t, err)

	acc, err = r.FindById(account.AccountID)
	assert.Nil(t, err)
	assert.EqualValues(t, 100, acc.Balance)
	assert.EqualValues(t, "0.03", fmt.Sprintf("%.2f", acc.Interest))

}

func TestPersonalAccountRepository_AllocateInterest(t *testing.T) {
	r := NewEntPersonalAccountRepository()
	jr := NewEntJobRepository()

	// TEST successful allocation
	jobs, err := jr.FindAll(nil, nil)
	assert.Len(t, jobs, 0)

	account, err := createPersonalAccount(datastore.EntClient)
	assert.Nil(t, err)

	err = r.AllocateInterest()
	assert.Nil(t, err)

	jobs, err = jr.FindAll(nil, nil)
	assert.Nil(t, err)
	assert.Len(t, jobs, 1)
	assert.EqualValues(t, 0, jobs[0].TotalProcessed)
	assert.EqualValues(t, "COMPLETED", jobs[0].Status)

	// TEST does not reallocate if completed
	err = r.AllocateInterest()
	assert.Error(t, err, internal_errors.JobAlreadyCompleted)

	jobs2, err := jr.FindAll(nil, nil)
	assert.Nil(t, err)
	assert.Len(t, jobs, 1)
	assert.EqualValues(t, 0, jobs[0].TotalProcessed)
	assert.EqualValues(t, "COMPLETED", jobs[0].Status)
	assert.EqualValues(t, jobs[0].UpdatedAt, jobs2[0].UpdatedAt)

	// TEST Interest allocation
	TruncateTable("jobs")
	viper.Set("APR", 12)

	acc, err := r.FindById(account.ID)
	assert.Nil(t, err)
	assert.EqualValues(t, 0, acc.Balance)
	assert.EqualValues(t, 0, acc.Interest)

	acc, err = acc.Update().AddBalance(100).AddInterest(10).Save(context.Background())
	assert.Nil(t, err)

	err = r.AllocateInterest()
	assert.Nil(t, err)

	acc, err = r.FindById(account.ID)
	assert.Nil(t, err)
	assert.EqualValues(t, 110, acc.Balance)
	assert.EqualValues(t, 0, acc.Interest)

	// TEST Batching
	TruncateTable("jobs")
	TruncateTable("personal_account_transactions")
	TruncateTable("personal_accounts")

	acc, err = createPersonalAccount(datastore.EntClient)
	assert.Nil(t, err)

	acc2, err := createPersonalAccountRandom(datastore.EntClient)
	assert.Nil(t, err)

	acc, err = acc.Update().AddBalance(100).AddInterest(10).Save(context.Background())
	assert.Nil(t, err)

	acc2, err = acc2.Update().AddBalance(1000).AddInterest(100).Save(context.Background())
	assert.Nil(t, err)

	_, err = datastore.EntClient.Job.Create().
		SetName("INTEREST_ALLOCATION").
		SetDate(time.Now().Format("2006-01")).
		SetBatch(1).
		Save(context.Background())
	assert.Nil(t, err)

	err = r.AllocateInterest()
	assert.Nil(t, err)

	acc, err = r.FindById(account.ID)
	assert.Nil(t, err)
	assert.EqualValues(t, 110, acc.Balance)
	assert.EqualValues(t, 0, acc.Interest)

	acc2, err = r.FindById(acc2.ID)
	assert.Nil(t, err)
	assert.EqualValues(t, 1100, acc2.Balance)
	assert.EqualValues(t, 0, acc2.Interest)

	// TEST Batching 2
	TruncateTable("jobs")
	TruncateTable("personal_account_transactions")
	TruncateTable("personal_accounts")

	acc, err = createPersonalAccount(datastore.EntClient)
	assert.Nil(t, err)

	acc2, err = createPersonalAccountRandom(datastore.EntClient)
	assert.Nil(t, err)

	acc, err = acc.Update().AddBalance(100).AddInterest(10).Save(context.Background())
	assert.Nil(t, err)

	acc2, err = acc2.Update().AddBalance(1000).AddInterest(100).Save(context.Background())
	assert.Nil(t, err)

	_, err = datastore.EntClient.Job.Create().
		SetName("INTEREST_ALLOCATION").
		SetDate(time.Now().Format("2006-01")).
		SetBatch(1).
		SetLastProcessedID(1).
		Save(context.Background())
	assert.Nil(t, err)

	err = r.AllocateInterest()
	assert.Nil(t, err)

	// Not allocated
	acc, err = r.FindById(acc.ID)
	assert.Nil(t, err)
	assert.EqualValues(t, 100, acc.Balance)
	assert.EqualValues(t, 10, acc.Interest)

	// Allocated
	acc2, err = r.FindById(acc2.ID)
	assert.Nil(t, err)
	assert.EqualValues(t, 1100, acc2.Balance)
	assert.EqualValues(t, 0, acc2.Interest)

	jobs, err = jr.FindAll(nil, nil)
	assert.Nil(t, err)
	assert.Len(t, jobs, 1)
	assert.EqualValues(t, 1, jobs[0].TotalProcessed)
	assert.EqualValues(t, 2, jobs[0].LastProcessedID)
	assert.EqualValues(t, "COMPLETED", jobs[0].Status)

	// TEST Batching - Heavy
	TruncateTable("jobs")
	TruncateTable("personal_account_transactions")
	TruncateTable("personal_accounts")

	for i := 0; i < 10000; i++ {
		_, err := datastore.EntClient.PersonalAccount.Create().
			SetAccountID(uint64(utils.RandomInt(1, 1000))).
			SetType(utils.RandomString(6)).
			SetBalance(float32(utils.RandomInt(0, 1000))).
			SetInterest(float32(utils.RandomInt(0, 100))).
			Save(context.Background())
		assert.Nil(t, err)
	}

	err = r.AllocateInterest()
	assert.Nil(t, err)

	jobs, err = jr.FindAll(nil, nil)
	assert.Nil(t, err)
	assert.Len(t, jobs, 1)
	assert.LessOrEqual(t, jobs[0].TotalProcessed, uint(10000))
	assert.EqualValues(t, 10000, jobs[0].LastProcessedID)
	assert.EqualValues(t, "COMPLETED", jobs[0].Status)

}
