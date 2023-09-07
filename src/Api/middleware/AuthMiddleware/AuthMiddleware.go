package AuthMiddleware

import (
	"go-learning/src/Utils/Jwt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func excludePaths(path string, arrPath []string) bool {
	for _, pattern := range arrPath {
		if strings.HasSuffix(pattern, "*") && strings.HasPrefix(path, strings.TrimSuffix(pattern, "*")) {
			return true
		} else if path == pattern {
			return true
		}
	}
	return false
}

var AuthData any
var IgnoreAuth []string

func CheckAuth(c *fiber.Ctx) error {

	if excludePaths(c.Path(), IgnoreAuth) {
		return c.Next()
	} else {
		header := c.GetReqHeaders()
		token := header["Authorization"]
		data, status, err := Jwt.VerifyToken(string(token))
		AuthData = data

		if err != nil {
			log.Error(err)
		}

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
