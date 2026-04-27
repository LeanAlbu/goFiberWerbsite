package main

import (
	"log"

	"github.com/LeanAlbu/goFiberWebsite/internal/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)


func main(){
	app := fiber.New(fiber.Config{
		AppName: "Teste",
	})

	app.Use(logger.New())
	app.Use(recover.New())

	routes.SetupRegistry(app)

	log.Println("Iniciando servidor na porta 3000...")
	log.Fatal(app.Listen(":3000"))
}
