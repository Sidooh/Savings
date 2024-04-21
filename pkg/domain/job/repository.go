package domain

import (
	"Savings/ent"
	"Savings/pkg/repositories/filters"
	"Savings/utils"
)

type JobRepository interface {
	FindAll(*utils.Paginator, *filters.JobFilters) (ent.Jobs, error)
	//FindById(uint64) (*ent.Job, error)
}
