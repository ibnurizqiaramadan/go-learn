package QueryMutation

import (
	"context"
	"go-learning/src/Utils/GraphqlClient"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/hasura/go-graphql-client"
)

type users_insert_input map[string]interface{}

func MutationUsers(c *fiber.Ctx) error {
	var q struct {
		InsertUsers struct {
			Returning []struct {
				ID              any `graphql:"id"`
				Email           any `graphql:"email"`
				Username        any `graphql:"username"`
				Fullname        any `graphql:"fullname"`
				Encrypt         any `graphql:"encrypt"`
				IsEmailVerified any `graphql:"is_email_verified"`
			} `graphql:"returning"`
		} `graphql:"insert_users(objects: $objects)"`
	}

	variables := map[string]interface{}{
		"objects": []users_insert_input{
			{
				"fullname":          "psp1",
				"email":             "psp1@gmail.com",
				"username":          "psp1",
				"encrypt":           "psp1",
				"is_email_verified": true,
			},
		},
	}

	client := GraphqlClient.CreateAdmin()
	err := client.Mutate(context.Background(), &q, variables, graphql.OperationName("InsertUsers"))

	messages := "" // Declare messages here with an initial empty value

	if err != nil {
		log.Debug(err)
		log.Debug(variables)
		log.Debug(q)
		return c.Status(fiber.StatusBadGateway).SendString("Something went wrong : " + err.Error())
	} else {
		if len(q.InsertUsers.Returning) == 0 {
			messages = "fail2" // Update messages here
		} else {
			messages = "success" // Update messages here
		}
	}

	return c.JSON(fiber.Map{
		"statusCode": fiber.StatusOK,
		"data": fiber.Map{
			"valid":    true,
			"messages": messages, // Use the updated messages variable here
		},
	})
}

func MutationUpdateUsers(c *fiber.Ctx) error {
	var q struct {
		UpdateUsers struct {
			Returning []struct {
				ID any `graphql:"id"`
				// Add other fields as needed for the update
			} `graphql:"returning"`
		} `graphql:"update_users(where: {id: {_eq: \"32d35924-3377-46ca-9971-88a8b77683ba\"}}, _set: {fullname: \"psp10\"})"`
	}

	// updateData := map[string]interface{}{
	// 	"fullname": "psp10",
	// 	// Add other fields as needed for the update
	// }

	variables := map[string]interface{}{}

	log.Debug(q)
	client := GraphqlClient.CreateAdmin()
	err := client.Mutate(context.Background(), &q, variables, graphql.OperationName("UpdateUsers"))

	messages := ""

	if err != nil {
		log.Debug(err)
		log.Debug(q)
		return c.Status(fiber.StatusBadGateway).SendString("Something went wrong: " + err.Error())
	} else {
		if len(q.UpdateUsers.Returning) == 0 {
			messages = "fail2"
		} else {
			messages = "success"
		}
	}

	return c.JSON(fiber.Map{
		"statusCode": fiber.StatusOK,
		"data": fiber.Map{
			"valid":    true,
			"messages": messages,
		},
	})
}
