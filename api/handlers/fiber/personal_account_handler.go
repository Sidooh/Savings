package fiberHandlers

import (
	domain "Savings/pkg/domain/personal_account"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type PersonalAccountHandler interface {
	Get(c *fiber.Ctx) error
	//GetById(c  *fiber.Ctx)
	//Create(c  *fiber.Ctx)
}

type personalAccountHandler struct {
	personalAccountService domain.PersonalAccountService
}

func NewPersonalAccountHandler(service domain.PersonalAccountService) PersonalAccountHandler {
	return &personalAccountHandler{personalAccountService: service}
}

func (h *personalAccountHandler) Get(c *fiber.Ctx) error {
	if personalAccounts, err := h.personalAccountService.FindAllPersonalAccounts(); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	} else {
		return c.Status(http.StatusOK).JSON(personalAccounts)
	}
}
