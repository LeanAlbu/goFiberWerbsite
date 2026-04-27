package routes

import (
	"github.com/LeanAlbu/goFiberWebsite/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRegistry(app *fiber.App){
	api := app.Group("api/v1")

	snippetsRoutes := api.Group("/snippets")
	snippetsRoutes.Get("/", handlers.GetSnippet)
	snippetsRoutes.Post("/", handlers.CreateSnippet)
	
}
