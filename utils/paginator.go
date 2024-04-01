package utils

import "github.com/gofiber/fiber/v2"

type Paginator struct {
	page     int
	pageSize int
	sort     string
	dir      string
}

// Use getters in case we need to add validation of values later
func (p *Paginator) Page() int {
	return p.page
}

func (p *Paginator) PageSize() int {
	if p == nil {
		return 15
	}
	return p.pageSize
}

func (p *Paginator) Offset() int {
	if p == nil {
		return 0
	}
	return p.pageSize * (p.page - 1)
}

func (p *Paginator) Meta() map[string]int {
	return map[string]int{
		"page":      p.page,
		"page_size": p.pageSize,
	}
}

func PaginatorFromFiber(c *fiber.Ctx) *Paginator {
	page := c.QueryInt("page", 1)
	pageSize := c.QueryInt("page_size", 15)

	if pageSize < 15 {
		pageSize = 15
	}

	if page < 1 {
		page = 1
	}

	return &Paginator{
		page:     page,
		pageSize: pageSize,
	}
}
