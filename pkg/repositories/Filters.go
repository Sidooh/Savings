package repositories

import (
	"github.com/gofiber/fiber/v2"
)

type PersonalAccountFilters struct {
	AccountId uint64
}

func PersonalAccountFiltersFromFiber(c *fiber.Ctx) *PersonalAccountFilters {
	accountId := c.QueryInt("account_id")

	return &PersonalAccountFilters{
		AccountId: uint64(accountId),
	}
}
