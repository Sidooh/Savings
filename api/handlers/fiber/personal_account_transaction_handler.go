package fiberHandlers

import (
	domain "Savings/pkg/domain/personal_account_transaction"
	"Savings/pkg/repositories/filters"
	"Savings/utils"
	internal_errors "Savings/utils/errors"
	"Savings/utils/responses"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type PersonalAccountTransactionHandler interface {
	Get(c *fiber.Ctx) error
	GetById(c *fiber.Ctx) error
}

type personalAccountTransactionHandler struct {
	personalAccountTransactionService domain.PersonalAccountTransactionService
}

func NewPersonalAccountTransactionHandler(service domain.PersonalAccountTransactionService) PersonalAccountTransactionHandler {
	return &personalAccountTransactionHandler{personalAccountTransactionService: service}
}

func (h *personalAccountTransactionHandler) Get(c *fiber.Ctx) error {
	paginator := utils.PaginatorFromFiber(c)
	filters := filters.PersonalAccountTransactionFiltersFromFiber(c)
	if personalAccountTransactions, err := h.personalAccountTransactionService.FindAllPersonalAccountTransactions(paginator, filters); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.ErrorResponse(err, c))
	} else {
		return c.Status(http.StatusOK).JSON(responses.PaginatedResponse(personalAccountTransactions, paginator))
	}
}

func (h *personalAccountTransactionHandler) GetById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.JSON(responses.ErrorResponse(internal_errors.InvalidIdParameter, c))
	}

	if personalAccountTransaction, err := h.personalAccountTransactionService.FindPersonalAccountTransactionById(uint64(id)); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.ErrorResponse(err, c))
	} else {
		return c.Status(http.StatusOK).JSON(responses.SuccessResponse(personalAccountTransaction))
	}
}
