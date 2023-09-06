package Billings

import (
	"go-learning/src/Utils/Validation"

	"github.com/gofiber/fiber/v2"
	"github.com/stripe/stripe-go/customer"
)

type GetCustomerId struct {
	Customer_id string `json:"customer_id" validate:"required"`
}

func GetCustomer(c *fiber.Ctx) error {
	custommer := new(GetCustomerId)
	errors, isValid := Validation.ValidateInput(c, custommer)
	if !isValid {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"statusCode": fiber.StatusBadRequest, "messages": "Invalid Input", "errors": errors})
	}
	cus, _ := customer.Get(custommer.Customer_id, nil)

	response := fiber.Map{
		"statusCode": fiber.StatusOK,
		"data": fiber.Map{
			"valid":       true,
			"messages":    "success-get-data-customer",
			"customer_id": cus.ID,
			"email":       cus.Email,
			"name":        cus.Name,
			"description": cus.Description,
		},
	}

	return c.JSON(response)
}
