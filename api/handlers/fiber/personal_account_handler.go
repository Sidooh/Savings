package fiberHandlers

import (
	"Savings/api/middleware"
	"Savings/ent"
	domain "Savings/pkg/domain/personal_account"
	"Savings/utils/responses"
	"errors"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type PersonalAccountHandler interface {
	Get(c *fiber.Ctx) error
	GetById(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
}

type CreatePersonalAccountRequest struct {
	AccountId uint64 `json:"account_id" validate:"required,numeric"`
	Type      string `json:"type" validate:"required"`
}

type personalAccountHandler struct {
	personalAccountService domain.PersonalAccountService
}

func NewPersonalAccountHandler(service domain.PersonalAccountService) PersonalAccountHandler {
	return &personalAccountHandler{personalAccountService: service}
}

func (h *personalAccountHandler) Get(c *fiber.Ctx) error {
	if personalAccounts, err := h.personalAccountService.FindAllPersonalAccounts(); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.ErrorResponse(err))
	} else {
		return c.Status(http.StatusOK).JSON(responses.SuccessResponse(personalAccounts))
	}
}

func (h *personalAccountHandler) GetById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		c.Status(http.StatusBadRequest)
		return c.JSON(responses.ErrorResponse(errors.New("invalid id parameter")))
	}

	if personalAccount, err := h.personalAccountService.FindPersonalAccountById(uint64(id)); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.ErrorResponse(err))
	} else {
		return c.Status(http.StatusOK).JSON(responses.SuccessResponse(personalAccount))
	}
}

func (h *personalAccountHandler) Create(c *fiber.Ctx) error {
	var request CreatePersonalAccountRequest
	if err := middleware.BindAndValidateRequest(c, &request); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(responses.ErrorResponse(err))
	}

	json := ent.PersonalAccount{
		AccountID: request.AccountId,
		Type:      request.Type,
	}

	if personalAccount, err := h.personalAccountService.CreateUser(&json); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.ErrorResponse(err))
	} else {
		return c.Status(http.StatusCreated).JSON(responses.SuccessResponse(personalAccount))
	}
}
