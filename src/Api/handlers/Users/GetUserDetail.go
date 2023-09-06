package Users

import (
	"go-learning/src/Utils/Jwt"
	"go-learning/src/Utils/MysqlClient"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type UserDetail struct {
	Subscribtion int
	First_name   string
	Last_name    string
	Email        string
	Picture      string
	Gender       string
	Country      any
}

func GetUserDetails(c *fiber.Ctx) error {
	token := c.Request().Header.Peek("Authorization")
	// verify token
	_, status, err := Jwt.VerifyToken(string(token))
	if !status {
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
	}

	storage := MysqlClient.DatabaseMod()
	id := c.Params("id")
	query := "SELECT subscribtion, first_name, last_name, email, picture, gender, country FROM users WHERE email = ?"
	rows, err := storage.Conn().Query(query, id)
	if err != nil {
		log.Error(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error retrieving data")
	}
	defer rows.Close()

	// Fetch rows
	var user UserDetail
	for rows.Next() {
		err := rows.Scan(&user.Subscribtion, &user.First_name, &user.Last_name, &user.Email, &user.Picture, &user.Gender, &user.Country)
		if err != nil {
			log.Error(err)
			return c.Status(fiber.StatusInternalServerError).SendString("Error retrieving data")
		}
	}
	// check if user is empty
	if (user == UserDetail{}) {
		var response = map[string]interface{}{
			"message": "User not found",
			"data":    []User{},
			"status":  "failed",
		}
		return c.Status(fiber.StatusNotFound).JSON(response)
	}

	var response = fiber.Map{
		"message": "Successfully get user by id",
		"data":    user,
		"status":  "success",
	}
	return c.Status(fiber.StatusOK).JSON(response)
}
