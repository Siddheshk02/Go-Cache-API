package routes

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
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

func GetPosts(c *fiber.Ctx) error {
	id := c.Params("id") // Get the post ID from the request URL parameters
	if id == "" {
		log.Fatal("Invalid ID")
	}

	// Fetch the post data from the API
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/" + id)
	if err != nil {
		return c.JSON(fiber.Map{"error": err.Error()})
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return c.JSON(fiber.Map{"error": err.Error()})
	}

	var data Post
	err = json.Unmarshal(body, &data)
	if err != nil {
		return c.JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(data)

}

func GetTodos(c *fiber.Ctx) error {
	id := c.Params("id") // Get the post ID from the request URL parameters
	if id == "" {
		log.Fatal("Invalid ID")
	}

	// Fetch the todo data from the API
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/" + id)
	if err != nil {
		return c.JSON(fiber.Map{"error": err.Error()})
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return c.JSON(fiber.Map{"error": err.Error()})
	}

	var data Todo
	err = json.Unmarshal(body, &data)
	if err != nil {
		return c.JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(data)

}
