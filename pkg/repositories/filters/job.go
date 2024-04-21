package filters

import (
	"github.com/gofiber/fiber/v2"
)

type JobFilters struct {
	Status string
}

func JobFiltersFromFiber(c *fiber.Ctx) *JobFilters {
	status := c.Query("status")

	return &JobFilters{
		Status: status,
	}
}
