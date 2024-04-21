package entRepo

import (
	"Savings/ent"
	"Savings/ent/job"
	"Savings/ent/personalaccount"
	"Savings/pkg/datastore"
	domain "Savings/pkg/domain/personal_account"
	"Savings/pkg/repositories/filters"
	"Savings/utils"
	internal_errors "Savings/utils/errors"
	"Savings/utils/logger"
	"context"
	"entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/spf13/viper"
	"time"
)

type personalAccountRepository struct {
	client   *ent.PersonalAccountClient
	txClient *ent.PersonalAccountTransactionClient
}

func (p *personalAccountRepository) FindAll(paginator *utils.Paginator, filters *filters.PersonalAccountFilters) (accounts ent.PersonalAccounts, err error) {
	q := p.client.Query()

	if filters != nil && filters.AccountId != 0 {
		q = q.Where(personalaccount.AccountID(filters.AccountId))
	}

	return q.Limit(paginator.PageSize()).Offset(paginator.Offset()).All(context.Background())
}

func (p *personalAccountRepository) FindById(id uint64) (account *ent.PersonalAccount, err error) {
	return p.client.Query().Where(personalaccount.ID(id)).First(context.Background())
}

func (p *personalAccountRepository) Create(data *ent.PersonalAccount) (account *ent.PersonalAccount, err error) {
	return p.client.Create().SetAccountID(data.AccountID).SetType(data.Type).Save(context.Background())
}

func (p *personalAccountRepository) Credit(id uint64, amount float32, description string) error {

	if err := WithTx(context.Background(), datastore.EntClient, func(tx *ent.Tx) error {
		account, err := tx.PersonalAccount.Query().Where(personalaccount.ID(id)).First(context.Background())
		if err != nil {
			return err
		}

		balance := account.Balance + amount

		_, err = tx.PersonalAccountTransaction.Create().
			SetType("CREDIT").
			SetPersonalAccountID(id).
			SetAmount(amount).
			SetBalance(balance).
			SetDescription(description).
			SetStatus("COMPLETED").
			Save(context.Background())
		if err != nil {
			return err
		}

		account.Balance = balance
		_, err = account.Update().SetBalance(balance).Save(context.Background())
		if err != nil {
			return err
		}

		return nil

	}); err != nil {
		return err
	}

	return nil

}

func (p *personalAccountRepository) Debit(id uint64, amount float32, description string) error {

	if err := WithTx(context.Background(), datastore.EntClient, func(tx *ent.Tx) error {
		account, err := tx.PersonalAccount.Query().Where(personalaccount.ID(id)).First(context.Background())
		if err != nil {
			return err
		}

		if account.Balance < amount {
			return internal_errors.InsufficientBalance
		}

		balance := account.Balance - amount

		_, err = tx.PersonalAccountTransaction.Create().
			SetType("DEBIT").
			SetPersonalAccountID(id).
			SetAmount(amount).
			SetBalance(balance).
			SetDescription(description).
			SetStatus("COMPLETED").
			Save(context.Background())
		if err != nil {
			return err
		}

		account.Balance = balance
		_, err = account.Update().SetBalance(balance).Save(context.Background())
		if err != nil {
			return err
		}

		return nil

	}); err != nil {
		return err
	}

	return nil

}

func (p *personalAccountRepository) CalculateInterest() error {
	dailyInterest := utils.GetDailyInterestRate()
	if dailyInterest == 0 {
		return internal_errors.DailyRateError
	}

	date := time.Now().Format("2006-01-02")
	job, err := datastore.EntClient.Job.Query().Where(job.Date(date)).First(context.Background())
	if err != nil {
		job, err = datastore.EntClient.Job.Create().
			SetName("INTEREST_CALCULATION").
			SetDate(date).
			SetData(map[string]interface{}{"rate": dailyInterest}).
			Save(context.Background())

		if err != nil {
			logger.Log.Error(err.Error())
			return internal_errors.JobCreationFailed
		}

		// TODO: Batch below

		//batch := 100
		//
		//accountIds, err := p.client.Query().
		//	Select(personalaccount.FieldID).
		//	Order(personalaccount.ByID(sql.OrderDesc())).
		//	Limit(batch).
		//	All(context.Background())
		//fmt.Println(accountIds, err)
		//if err != nil {
		//	return err
		//}
		//
		//if len(accountIds) == 0 {
		//	_, err := job.Update().SetStatus("COMPLETED").Save(context.Background())
		//	if err != nil {
		//		return err
		//	}
		//}
		//
		//// Update all entities by incrementing "interest" by new interest
		//updateQuery := fmt.Sprintf(`UPDATE personal_accounts SET interest = interest + balance * %v WHERE id BETWEEN %v AND %v;`, dailyInterest, 1, batch)
		//
		//res, err := datastore.EntClient.ExecContext(context.Background(), updateQuery)
		//fmt.Println(res.LastInsertId())
		//fmt.Println(res.RowsAffected())
		////if err != nil {
		////	// handle error
		////}
		//
		//fmt.Println(accountIds)

		// TODO: Batch above

		// With few accounts (< 10000) this should be fine
		updateQuery := fmt.Sprintf(`UPDATE personal_accounts SET interest = interest + balance * %v;`, dailyInterest)

		res, err := datastore.EntClient.ExecContext(context.Background(), updateQuery)
		if err != nil {
			logger.Log.Error(err.Error())
			return internal_errors.JobProcessingFailed
		}

		affected, _ := res.RowsAffected()

		_, err = job.Update().SetStatus("COMPLETED").SetTotalProcessed(uint(affected)).Save(context.Background())
		if err != nil {
			logger.Log.Error(err.Error())
			return internal_errors.JobProcessingFailed
		}

		return err
	}

	if job.Status == "COMPLETED" {
		return internal_errors.JobAlreadyCompleted
	}

	return nil
}

func (p *personalAccountRepository) AllocateInterest() error {
	month := time.Now().Format("2006-01")
	job, err := datastore.EntClient.Job.Query().Where(job.Date(month)).First(context.Background())
	if err != nil {
		job, err = datastore.EntClient.Job.Create().
			SetName("INTEREST_ALLOCATION").
			SetDate(month).
			Save(context.Background())

		if err != nil {
			logger.Log.Error(err.Error())
			return internal_errors.JobCreationFailed
		}

		return p.batchProcess(job)
	}

	return p.batchProcess(job)

	return nil
}

func (p *personalAccountRepository) batchProcess(job *ent.Job) error {
	if job.Status == "COMPLETED" {
		return internal_errors.JobAlreadyCompleted
	}

	if job.Status == "PENDING" || job.Status == "PROCESSING" {
		accounts, err := p.client.Query().
			Select(personalaccount.FieldID, personalaccount.FieldBalance, personalaccount.FieldInterest).
			Order(personalaccount.ByID(sql.OrderAsc())).
			Where(personalaccount.IDGT(job.LastProcessedID), personalaccount.InterestGT(0)).
			Limit(job.Batch).
			All(context.Background())

		if len(accounts) == 0 {
			_, err = job.Update().SetStatus("COMPLETED").Save(context.Background())
			if err != nil {
				logger.Log.Error(err.Error())
				return internal_errors.JobProcessingFailed
			}

			return nil
		}

		_, err = p.txClient.
			MapCreateBulk(accounts, func(c *ent.PersonalAccountTransactionCreate, i int) {
				c.SetAmount(accounts[i].Interest).
					SetAccountID(accounts[i].ID).
					SetType("CREDIT").
					SetDescription("INTEREST").
					SetStatus("COMPLETED")
			}).
			Save(context.Background())

		updateQuery := fmt.Sprintf(
			`UPDATE personal_accounts SET balance = balance + interest, interest = 0 WHERE id >= %v AND id <= %v AND interest > 0;`,
			accounts[0].ID,
			accounts[len(accounts)-1].ID,
		)

		_, err = datastore.EntClient.ExecContext(context.Background(), updateQuery)
		if err != nil {
			logger.Log.Error(err.Error())
			return internal_errors.JobProcessingFailed
		}

		job, err = job.Update().
			SetStatus("PROCESSING").
			SetLastProcessedID(accounts[len(accounts)-1].ID).
			AddTotalProcessed(len(accounts)).
			Save(context.Background())
		if err != nil {
			logger.Log.Error(err.Error())
			return internal_errors.JobProcessingFailed
		}

	}

	return p.batchProcess(job)
}

func NewEntPersonalAccountRepository() domain.PersonalAccountRepository {
	c := datastore.EntClient
	if viper.GetBool("APP_DEBUG") {
		c = c.Debug()
	}

	return &personalAccountRepository{
		client:   c.PersonalAccount,
		txClient: c.PersonalAccountTransaction,
	}
}
