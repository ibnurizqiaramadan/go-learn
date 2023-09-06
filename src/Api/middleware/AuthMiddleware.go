package AuthMiddleware

import (
	"github.com/gofiber/fiber/v2"
)

func CheckAuth(c *fiber.Ctx) error {
	if c.Path() == "/" {
		return c.Next()
	} else {
		header := c.GetReqHeaders()
		token := header["Authorization"]
		// log.Info("token : " + token)
		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"statusCode": fiber.StatusUnauthorized,
				"data": fiber.Map{
					"message": "Unauthorized",
				},
			})
		}
		return c.Next()
	}
}
