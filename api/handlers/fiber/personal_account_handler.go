package fiberHandlers

import (
	"Savings/api/middleware"
	"Savings/ent"
	domain "Savings/pkg/domain/personal_account"
	"Savings/pkg/repositories/filters"
	"Savings/utils"
	internal_errors "Savings/utils/errors"
	"Savings/utils/responses"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type PersonalAccountHandler interface {
	Get(c *fiber.Ctx) error
	GetById(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error

	Deposit(c *fiber.Ctx) error
	Withdraw(c *fiber.Ctx) error
}

type CreatePersonalAccountRequest struct {
	AccountId uint64 `json:"account_id" validate:"required,numeric"`
	Type      string `json:"type" validate:"required"`
}

type CreditDebitPersonalAccountRequest struct {
	Amount float32 `json:"amount" validate:"required,numeric"`
}

type personalAccountHandler struct {
	personalAccountService domain.PersonalAccountService
}

func NewPersonalAccountHandler(service domain.PersonalAccountService) PersonalAccountHandler {
	return &personalAccountHandler{personalAccountService: service}
}

func (h *personalAccountHandler) Get(c *fiber.Ctx) error {
	paginator := utils.PaginatorFromFiber(c)
	filters := filters.PersonalAccountFiltersFromFiber(c)
	if personalAccounts, err := h.personalAccountService.FindAllPersonalAccounts(paginator, filters); err != nil {
		return c.JSON(responses.ErrorResponse(err, c))
	} else {
		return c.Status(http.StatusOK).JSON(responses.PaginatedResponse(personalAccounts, paginator))
	}
}

func (h *personalAccountHandler) GetById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.JSON(responses.ErrorResponse(internal_errors.InvalidIdParameter, c))
	}

	if personalAccount, err := h.personalAccountService.FindPersonalAccountById(uint64(id)); err != nil {
		return c.JSON(responses.ErrorResponse(err, c))
	} else {
		return c.Status(http.StatusOK).JSON(responses.SuccessResponse(personalAccount))
	}
}

func (h *personalAccountHandler) Create(c *fiber.Ctx) error {
	var request CreatePersonalAccountRequest
	if err := middleware.BindAndValidateRequest(c, &request); err != nil {
		return c.JSON(responses.ErrorResponse(err, c))
	}

	json := ent.PersonalAccount{
		AccountID: request.AccountId,
		Type:      request.Type,
	}

	if personalAccount, err := h.personalAccountService.CreatePersonalAccount(&json); err != nil {
		return c.JSON(responses.ErrorResponse(err, c))
	} else {
		return c.Status(http.StatusCreated).JSON(responses.SuccessResponse(personalAccount))
	}
}

func (h *personalAccountHandler) Deposit(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.JSON(responses.ErrorResponse(internal_errors.InvalidIdParameter, c))
	}

	var request CreditDebitPersonalAccountRequest
	if err = middleware.BindAndValidateRequest(c, &request); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(responses.ErrorResponse(err, c))
	}

	if err = h.personalAccountService.CreditPersonalAccount(uint64(id), request.Amount, "Deposit"); err != nil {
		return c.JSON(responses.ErrorResponse(err, c))
	} else {
		return c.Status(http.StatusCreated).JSON(responses.SuccessResponse(nil))
	}
}

func (h *personalAccountHandler) Withdraw(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.JSON(responses.ErrorResponse(internal_errors.InvalidIdParameter, c))
	}

	var request CreditDebitPersonalAccountRequest
	if err = middleware.BindAndValidateRequest(c, &request); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(responses.ErrorResponse(err, c))
	}

	if err = h.personalAccountService.DebitPersonalAccount(uint64(id), request.Amount, "Withdraw"); err != nil {
		return c.JSON(responses.ErrorResponse(err, c))
	} else {
		return c.Status(http.StatusCreated).JSON(responses.SuccessResponse(nil))
	}
}
