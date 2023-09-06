package Billings

import (
	"go-learning/src/Interfaces"
	"go-learning/src/Utils/StripeClient"
	"go-learning/src/Utils/Validation"

	"github.com/gofiber/fiber/v2"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/customer"
)

func AddCustomer(c *fiber.Ctx) error {

	newCustommer := Interfaces.AddCustomer{}

	error, isValid := Validation.ValidateInput(c, newCustommer)
	if !isValid {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"statusCode": fiber.StatusBadRequest, "messages": "Invalid Input", "errors": error})
	}

	StripeClient.InitStripe()
	params := &stripe.CustomerParams{
		Email:       stripe.String(newCustommer.Email),
		Name:        stripe.String(newCustommer.Name),
		Description: stripe.String(newCustommer.Description),
	}
	cus, _ := customer.New(params)

	return c.JSON(fiber.Map{
		"statusCode": 200,
		"data": fiber.Map{
			"valid":       true,
			"messages":    "success-create-customer",
			"customer_id": cus.ID,
		},
	})
}
