package Users

import (
	"context"
	"go-learning/src/Utils/Jwt"
	"go-learning/src/Utils/MysqlClient"
	"go-learning/src/Utils/Validation"
	"time"

	"go-learning/src/Utils/RedisClient"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type User struct {
	Id        string `json:"id"`
	Username  string `json:"username" validate:"required"`
	Role      string `json:"role"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
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

	var response = fiber.Map{
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

func Login(c *fiber.Ctx) error{
	// Get data from body request
	dataBody := new(User)
	errors, isValid := Validation.ValidateInput(c, dataBody)
	if !isValid {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": errors})
	}
	// Get username body request
	username := dataBody.Username
	// Get data from database
	storage := MysqlClient.CreateMysqlClient()
	// create query
	query := ("SELECT id, username, role, createdAt, updatedAt FROM users WHERE username = ?")
	// execute query
	rows, err := storage.Conn().Query(query, username)
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

	// create last login
	lastlogin:= time.Now()
	key := "lastlogin_" + user.Username
	ctx := context.Background()

	// set to redis
	Client := RedisClient.Client
	errRedis := Client.Set(ctx,key,lastlogin, 30).Err()
	if errRedis != nil {
		log.Error(errRedis)
		return c.Status(fiber.StatusInternalServerError).SendString("Error set data to redis")
	}

	// check if user is empty
	if user.Id == "" {
		var response = fiber.Map{
			"message": "User not found",
			"data":    []User{},
			"status":  fiber.StatusNotFound,
		}
		return c.Status(fiber.StatusNotFound).JSON(response)
	}
	// create data playload
	data := Jwt.Claims{
		Authorized: true,
		User: user.Username,
	}
	// create token
	token:= Jwt.CreateToken(Jwt.Claims(data))
	// Return data
	var response = fiber.Map{
		"message": "Successfully get user by id",
		"status":  fiber.StatusOK,
		"token": token,
	}
	// Return data
	return c.Status(fiber.StatusOK).JSON(response)
}