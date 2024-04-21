package entRepo

import (
	"Savings/ent"
	"Savings/ent/job"
	"Savings/pkg/datastore"
	domain "Savings/pkg/domain/job"
	"Savings/pkg/repositories/filters"
	"Savings/utils"
	"context"
	"github.com/spf13/viper"
)

type jobRepository struct {
	client *ent.JobClient
}

func (j *jobRepository) FindAll(paginator *utils.Paginator, filters *filters.JobFilters) (jobs ent.Jobs, err error) {
	q := j.client.Query()

	if filters != nil && filters.Status != "" {
		q = q.Where(job.Status(filters.Status))
	}

	return q.Limit(paginator.PageSize()).Offset(paginator.Offset()).All(context.Background())
}

//func (j *jobRepository) FindById(id uint64) (job *ent.Job, err error) {
//	return j.client.Query().Where(job.ID(id)).First(context.Background())
//}

func NewEntJobRepository() domain.JobRepository {
	c := datastore.EntClient
	if viper.GetBool("APP_DEBUG") {
		c = c.Debug()
	}

	return &jobRepository{
		client: c.Job,
	}
}
