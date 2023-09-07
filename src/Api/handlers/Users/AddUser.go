package Users

import (
	AuthMiddleware "go-learning/src/Api/middleware"
	"go-learning/src/Utils/MysqlClient"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func AddUsers(c *fiber.Ctx) error {
	authData := AuthMiddleware.AuthData // get decoded auth token
	return c.Status(fiber.StatusOK).JSON(authData)
}

func AddUserMysql(c *fiber.Ctx) error {
	dataBody := new(User)
	if err := c.BodyParser(dataBody); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Error retrieving data")
	}

	id := uuid.New()
	username := dataBody.Username
	password := "123456"
	role := dataBody.Role
	createAt := time.Now()
	updateAt := time.Now()

	storage := MysqlClient.CreateMysqlClient()
	query := "INSERT INTO users (id, username, password, role, createdAt, updatedAt) VALUES (?, ?, ?, ?, ?, ?)"
	_, err := storage.Conn().Query(query, id, username, password, role, createAt, updateAt)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error retrieving data")
	}
	var response = map[string]interface{}{
		"status":  fiber.StatusOK,
		"message": "Successfully add user",
	}
	return c.Status(fiber.StatusOK).JSON(response)
}
