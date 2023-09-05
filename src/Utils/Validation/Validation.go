package Validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var Validate = validator.New()

type CustomError struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

func ValidateInput(c *fiber.Ctx, input interface{}) ([]CustomError, bool) {
	errors := []CustomError{}

	if err := c.BodyParser(input); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid JSON input"})
		return nil, false
	}

	if err := Validate.Struct(input); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, ve := range validationErrors {
				errors = append(errors, CustomError{
					Field: ve.Field(),
					Error: ve.Tag(),
				})
			}
		}
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": errors})
		return errors, false
	}

	return nil, true
}
