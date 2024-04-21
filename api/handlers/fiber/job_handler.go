package fiberHandlers

import (
	domain "Savings/pkg/domain/job"
	"Savings/pkg/repositories/filters"
	"Savings/utils"
	"Savings/utils/responses"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type JobHandler interface {
	Get(c *fiber.Ctx) error
	CalculateInterest(c *fiber.Ctx) error
	AllocateInterest(c *fiber.Ctx) error
}

type jobHandler struct {
	jobService domain.JobService
}

func (h *jobHandler) Get(c *fiber.Ctx) error {
	paginator := utils.PaginatorFromFiber(c)
	filters := filters.JobFiltersFromFiber(c)
	if jobs, err := h.jobService.FindAllJobs(paginator, filters); err != nil {
		return c.JSON(responses.ErrorResponse(err, c))
	} else {
		return c.Status(http.StatusOK).JSON(responses.PaginatedResponse(jobs, paginator))
	}
}

func (h *jobHandler) CalculateInterest(c *fiber.Ctx) error {
	if err := h.jobService.CalculateInterest(); err != nil {
		return c.JSON(responses.ErrorResponse(err, c))
	} else {
		return c.Status(http.StatusOK).JSON(responses.SuccessResponse(true))
	}
}

func (h *jobHandler) AllocateInterest(c *fiber.Ctx) error {
	if err := h.jobService.AllocateInterest(); err != nil {
		return c.JSON(responses.ErrorResponse(err, c))
	} else {
		return c.Status(http.StatusOK).JSON(responses.SuccessResponse(true))
	}
}

func NewJobHandler(service domain.JobService) JobHandler {
	return &jobHandler{jobService: service}
}
