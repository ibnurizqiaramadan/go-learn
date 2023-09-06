package AuthMiddleware

import (
	"go-learning/src/Utils/Jwt"

	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/log"
)

func CheckAuth(c *fiber.Ctx) error {
	if c.Path() == "/" {
		return c.Next()
	} else {
		header := c.GetReqHeaders()
		token := header["Authorization"]
		_, status := Jwt.VerifyToken(string(token))
		// data, status := Jwt.VerifyToken(string(token))

		// log.Info(data, status)

		if token == "" || !status {
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
