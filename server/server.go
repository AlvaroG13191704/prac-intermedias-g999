package main

import (
	"server/api-go-test/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

func main() {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	app := fiber.New()

	// route
	app.Get("/", routes.Hello)
	app.Get("/todos", routes.GetTodos(client))
	app.Post("/todos", routes.CreateTodo(client))
	app.Put("/todos/:id", routes.UpdateTodo(client))
	app.Delete("/todos/:id", routes.DeleteTodo(client))

	app.Listen(":3000")
}
