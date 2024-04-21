package domain

import (
	"Savings/ent"
	domain "Savings/pkg/domain/personal_account"
	"Savings/pkg/repositories/filters"
	"Savings/utils"
)

type JobService interface {
	FindAllJobs(*utils.Paginator, *filters.JobFilters) (ent.Jobs, error)
	//FindJobById(uint64) (*ent.Job, error)

	CalculateInterest() error
	AllocateInterest() error
}

type jobService struct {
	repo     JobRepository
	pAccRepo domain.PersonalAccountRepository
}

func (j *jobService) CalculateInterest() error {
	return j.pAccRepo.CalculateInterest()
}

func (j *jobService) AllocateInterest() error {
	return j.pAccRepo.AllocateInterest()
}

func (j *jobService) FindAllJobs(paginator *utils.Paginator, filters *filters.JobFilters) (ent.Jobs, error) {
	return j.repo.FindAll(paginator, filters)
}

//func (j *jobService) FindJobById(id uint64) (*ent.Job, error) {
//	return j.repo.FindById(id)
//}

func NewJobService(repository JobRepository, pAccRepository domain.PersonalAccountRepository) JobService {
	return &jobService{repo: repository, pAccRepo: pAccRepository}
}
