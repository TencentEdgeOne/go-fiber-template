package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// Todo 待办事项模型
type Todo struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"createdAt"`
}

var (
	todoMu    sync.RWMutex
	todoSeq   = 3
	todoStore = []Todo{
		{ID: 1, Title: "Deploy to EdgeOne", Completed: true, CreatedAt: time.Now().Add(-48 * time.Hour)},
		{ID: 2, Title: "Write Go handlers", Completed: true, CreatedAt: time.Now().Add(-24 * time.Hour)},
		{ID: 3, Title: "Try Fiber framework", Completed: false, CreatedAt: time.Now()},
	}
)

func main() {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	app.Use(recover.New())

	// Welcome
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Welcome to EdgeOne Fiber Demo!",
			"version": "1.0.0",
			"routes": []string{
				"GET  /            - this page",
				"GET  /health      - health check",
				"GET  /api/todos   - list todos",
				"POST /api/todos   - create todo",
				"GET  /api/todos/:id         - get todo",
				"PATCH /api/todos/:id/toggle  - toggle todo",
				"DELETE /api/todos/:id        - delete todo",
			},
		})
	})

	// Health
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":    "ok",
			"framework": "fiber",
		})
	})

	// Todo CRUD
	api := app.Group("/api")

	api.Get("/todos", func(c *fiber.Ctx) error {
		todoMu.RLock()
		defer todoMu.RUnlock()
		return c.JSON(fiber.Map{"data": todoStore, "total": len(todoStore)})
	})

	api.Post("/todos", func(c *fiber.Ctx) error {
		var req struct {
			Title string `json:"title"`
		}
		if err := c.BodyParser(&req); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
		}
		if req.Title == "" {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "title is required"})
		}
		todoMu.Lock()
		todoSeq++
		todo := Todo{ID: todoSeq, Title: req.Title, Completed: false, CreatedAt: time.Now()}
		todoStore = append(todoStore, todo)
		todoMu.Unlock()
		return c.Status(http.StatusCreated).JSON(fiber.Map{"data": todo})
	})

	api.Get("/todos/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
		}
		todoMu.RLock()
		defer todoMu.RUnlock()
		for _, t := range todoStore {
			if t.ID == id {
				return c.JSON(fiber.Map{"data": t})
			}
		}
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "todo not found"})
	})

	api.Patch("/todos/:id/toggle", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
		}
		todoMu.Lock()
		defer todoMu.Unlock()
		for i := range todoStore {
			if todoStore[i].ID == id {
				todoStore[i].Completed = !todoStore[i].Completed
				return c.JSON(fiber.Map{"data": todoStore[i]})
			}
		}
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "todo not found"})
	})

	api.Delete("/todos/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
		}
		todoMu.Lock()
		defer todoMu.Unlock()
		for i, t := range todoStore {
			if t.ID == id {
				todoStore = append(todoStore[:i], todoStore[i+1:]...)
				return c.JSON(fiber.Map{"message": "deleted"})
			}
		}
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "todo not found"})
	})

	port := "9000"
	fmt.Printf("Fiber server starting on :%s\n", port)
	log.Fatal(app.Listen(":" + port))
}
