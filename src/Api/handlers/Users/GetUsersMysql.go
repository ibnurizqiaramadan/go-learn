package Users

import (
	"go-learning/src/Utils/Jwt"
	"go-learning/src/Utils/MysqlClient"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type User struct {
	Id        string
	Username  string
	Role      string
	CreatedAt string
	UpdatedAt string
}

func GetUsersMysql(c *fiber.Ctx) error {
	storage := MysqlClient.CreateMysqlClient()
	query := "SELECT id, username, role, createdAt, updatedAt FROM users ORDER BY id DESC"
	rows, err := storage.Conn().Query(query)
	if err != nil {
		log.Error(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error retrieving data")
	}
	defer rows.Close()

	// Fetch rows
	var users [](User)
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Username, &user.Role, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			log.Error(err)
			return c.Status(fiber.StatusInternalServerError).SendString("Error retrieving data")
		}
		users = append(users, user)
	}
	// Return data
	var response = map[string]interface{}{
		"message": "Successfully get all users",
		"data":    users,
		"status":  "success",
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

func GetUserById(c *fiber.Ctx) error {
	storage := MysqlClient.CreateMysqlClient()
	id := c.Params("id")
	query := "SELECT id, username, role, createdAt, updatedAt FROM users WHERE id = ?"
	rows, err := storage.Conn().Query(query, id)
	if err != nil {
		log.Error(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error retrieving data")
	}
	defer rows.Close()

	// Fetch rows
	var user User
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Role, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			log.Error(err)
			return c.Status(fiber.StatusInternalServerError).SendString("Error retrieving data")
		}
	}
	// check if user is empty
	if user.Id == "" {
		var response = map[string]interface{}{
			"message": "User not found",
			"data":    []User{},
			"status":  "failed",
		}
		return c.Status(fiber.StatusNotFound).JSON(response)
	}

	var response = map[string]interface{}{
		"message": "Successfully get user by id",
		"data":    user,
		"status":  "success",
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

func TestConnection(c *fiber.Ctx) error {
	storage := MysqlClient.DatabaseMod()
	err := storage.Conn().Ping()
	if err != nil {
		log.Error(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error retrieving data")
	}
	return c.Status(fiber.StatusOK).SendString("Successfully connected to database")
}

func GenerateToken(c *fiber.Ctx) error{
	// Create token
	token:= Jwt.GenerateJwtToken()
	
	// verify token
	verifyToken, status := Jwt.VerifyJwtToken(token)

	if !status {
		res := map[string]interface{}{
			"message": "Token is not valid",
			"status":  status,
		}
		return c.Status(fiber.StatusUnauthorized).JSON(res)
	}
	var response = map[string]interface{}{
		"message": "Successfully generate token",
		"data":    token,
		"verify": verifyToken,
		"status":  status,
	}
	// Return data
	return c.Status(fiber.StatusOK).JSON(response)
}