package Users

import (
	"go-learning/src/Utils/MysqlClient"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func AddUsers(c *fiber.Ctx) error {
	return c.Status(200).SendString("Add Users !")
}

func AddUserMysql(c *fiber.Ctx) error {
	id := uuid.New()
	username := c.Params("username")
	password := c.Params("password")
	role := c.Params("role")
	createAt := time.Now()
	updateAt := time.Now()

	storage := MysqlClient.CreateMysqlClient()
	query := "INSERT INTO users (id, username, password, role, createdAt, updatedAt) VALUES (?, ?, ?, ?, ?, ?)"
	_, err := storage.Conn().Query(query, id, username, password, role, createAt, updateAt)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error retrieving data")
	}
	return c.Status(200).SendString("Add User Mysql !")
}
