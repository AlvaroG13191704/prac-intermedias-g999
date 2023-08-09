package routes

import (
	"context"
	"encoding/json"
	"server/api-go-test/model"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

// GetTodos is a function that returns all the todos in the database
func GetTodos(client *redis.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := context.Background()
		// get all the keys in the database
		keys, err := client.Keys(ctx, "*").Result()
		if err != nil {
			panic(err)
		}

		// get all the values in the database
		values, err := client.MGet(ctx, keys...).Result()
		if err != nil {
			panic(err)
		}

		// create a slice of todos
		todos := make([]model.Todo, len(values))

		// iterate through the values and unmarshal them into a todo
		for i, v := range values {
			json.Unmarshal([]byte(v.(string)), &todos[i])
		}

		// return status 200 OK and the todos
		return c.Status(200).JSON(todos)

	}
}

// CreateTodo is a function that creates a todo in the database
func CreateTodo(client *redis.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {

		ctx := context.Background()

		todo := new(model.Todo)

		if err := c.BodyParser(todo); err != nil {
			return c.Status(400).SendString(err.Error())
		}

		// marshal todo into a JSON string
		todoJSON, _ := json.Marshal(todo)

		// set the key-value pair in the database
		err := client.Set(ctx, todo.ID, todoJSON, 0).Err()
		if err != nil {
			panic(err)
		}

		// return status 201 Created and id of the todo
		return c.Status(201).JSON(todo.ID)
	}
}

// UpdateTodo is a function that updates a todo in the database
func UpdateTodo(client *redis.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := context.Background()
		// get the id from the url
		id := c.Params("id")

		todo := new(model.Todo)

		if err := c.BodyParser(todo); err != nil {
			return c.Status(400).SendString(err.Error())
		}

		// marshal todo into a JSON string
		todoJSON, _ := json.Marshal(todo)

		// update the key-value pair in the database
		err := client.Set(ctx, id, todoJSON, 0).Err()
		if err != nil {
			panic(err)
		}

		// return status 200 OK and the todo
		return c.Status(200).JSON(todo.ID)
	}
}

// DeleteTodo is a function that deletes a todo in the database
func DeleteTodo(client *redis.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {

		ctx := context.Background()

		// get the id from the url
		id := c.Params("id")
		// delete the key-value pair in the database
		err := client.Del(ctx, id).Err()
		if err != nil {
			panic(err)
		}

		// return status 200 OK and the todo
		return c.Status(200).JSON(id)
	}
}
