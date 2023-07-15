package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/gofiber/swagger"

	_ "go-backend/docs"
	"go-backend/handlers"
)

// @title Collaborative Text Editor's Go backend
// @version	0.0.1
// @description	This is the backend written in Go for the Collaborative Text Editor
func main() {
	const port int = 8080

	app := fiber.New(fiber.Config{
		AppName:           "Collaborative Text Editor's Go backend",
		EnablePrintRoutes: true,
	})

	// Setup middlewares
	app.Use(logger.New())

	// Root path operation
	// @Summary Root path operation, serving as a hello world endpoint.
	// Tags root
	// @Produce json
	// @Success 200
	// @Router / [get]
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"hello": "world",
		})
	})

	// Default swagger docs handler
	app.Get("/docs/*", swagger.HandlerDefault)

	app.Get("/users/:id?", handlers.GetUserByIdHandler)
	app.Post("/users/", handlers.CreateUser)

	app.Get("/documents", handlers.GetAllDocuments)
	app.Get("/documents/:id", handlers.GetDocumentByIdHandler)
	app.Post("/documents", handlers.CreateDocument)

	app.Get("/editing-sessions", handlers.GetAllEditingSessions)
	app.Post("/editing-sessions", handlers.CreateEditingSession)

	app.Put("/editing-sessions/:user_id/:document_id", handlers.UpdateEditingSession)
	app.Post("/editing-sessions/:user_id/:document_id", handlers.WriteInEditingSession)
	app.Delete("/editing-sessions/:user_id/:document_id/:number_of_chars",
		handlers.DeleteInEditingSession)

	log.Fatal(app.Listen(fmt.Sprintf("localhost:%d", port)))
}
