package controllers

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// ParsePageQueryParam ...
func ParsePageQueryParam(c *fiber.Ctx) (int, error) {
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		return 0, err
	}
	if page < 0 {
		return 0, errors.New("page must be an positive integer")
	}
	if page == 0 {
		return 1, nil
	}
	return page, nil
}

//ParseLimitQueryParam ...
func ParseLimitQueryParam(c *fiber.Ctx) (int, error) {
	limit, err := strconv.Atoi(c.Query("limit", "10"))
	if err != nil {
		return 0, err
	}
	if limit < 0 {
		return 0, errors.New("limit must be an positive integer")
	}
	if limit == 0 {
		return 10, nil
	}
	return limit, nil
}
