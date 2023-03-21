package middleware

import (
	"encoding/json"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/patrickmn/go-cache"
)

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func CacheMiddleware(cache *cache.Cache) fiber.Handler {
	return func(c *fiber.Ctx) error {

		cacheEnabled := true // set to true by default
		if c.OriginalURL() == "/todos/:id" {
			cacheEnabled = false // disable caching for todos
		}

		if cacheEnabled {
			if c.Method() != "GET" {
				// Only cache GET requests
				return c.Next()
			}

			// Generate a cache key from the request path and query parameters
			cacheKey := c.Path() + "?" + c.Params("id")

			// Check if the response is already in the cache
			if cached, found := cache.Get(cacheKey); found {
				return c.JSON(cached)
			}

		}

		// Otherwise, execute the route handler and cache the response
		err := c.Next()
		if err != nil {
			return err
		}

		if cacheEnabled {
			var data Post
			cacheKey := c.Path() + "?" + c.Params("id")

			body := c.Response().Body()
			err = json.Unmarshal(body, &data)
			if err != nil {
				return c.JSON(fiber.Map{"error": err.Error()})
			}

			// Cache the response for 10 minutes
			cache.Set(cacheKey, data, 10*time.Minute)

		}

		return nil
	}
}
