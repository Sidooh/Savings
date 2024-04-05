package filters

import (
	"github.com/gofiber/fiber/v2"
)

type PersonalAccountTransactionFilters struct {
	PersonalAccountId uint64
	Type              string
	Status            string
}

func PersonalAccountTransactionFiltersFromFiber(c *fiber.Ctx) *PersonalAccountTransactionFilters {
	personalAccountId := c.QueryInt("personal_account_id")

	return &PersonalAccountTransactionFilters{
		PersonalAccountId: uint64(personalAccountId),
	}
}
