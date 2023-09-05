package BasicUsages

import (
	"github.com/gofiber/fiber/v2"
)

func GetDatas(c *fiber.Ctx) error {
	// this getting param
	id := c.Query("id")

	username := c.Query("username")
	if username == "" {
		return c.JSON(fiber.Map{
			"statusCode": fiber.StatusBadRequest,
			"data": fiber.Map{
				"valid":    false,
				"messages": "username is required",
			},
		})
	}
	// this for getting args
	name := c.Params("name")

	return c.JSON(fiber.Map{
		"statusCode": fiber.StatusOK,
		"data": fiber.Map{
			"valid":    true,
			"messages": "success-get-data",
			"data": fiber.Map{
				"id":       id,
				"name":     name,
				"username": username,
			},
		},
	})
}
