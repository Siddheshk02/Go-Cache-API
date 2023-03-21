package main

import (
	"time"

	"github.com/Siddheshk02/Go-Cache-API/middleware"
	"github.com/Siddheshk02/Go-Cache-API/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/patrickmn/go-cache"
)

func main() {
	app := fiber.New()

	cache := cache.New(10*time.Minute, 20*time.Minute) // setting default expiration time and clearance time

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	// Define the routes for posts and todos

	app.Get("/posts/:id", middleware.CacheMiddleware(cache), routes.GetPosts)

	app.Get("/todos/:id", middleware.CacheMiddleware(cache), routes.GetTodos)

	app.Listen(":8000")
}
