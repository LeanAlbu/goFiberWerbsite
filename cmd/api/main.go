package main

import (
	"log"

	"github.com/LeanAlbu/goFiberWebsite/internal/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)


func main(){

	engine := html.New("./views", ".html")


	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", handlers.RenderHome)

	log.Println("Iniciando servidor na porta 3000...")
	log.Fatal(app.Listen(":3000"))
}
