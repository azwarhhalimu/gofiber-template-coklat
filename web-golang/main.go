package main

import (
	"belajar/web-golang/app/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/jet/v2"
)

func main() {

	template := jet.New("./app/views", ".jet")

	app := fiber.New(fiber.Config{
		Views: template,
	})
	app.Static("/storage", "./assets")
	public := app.Group("/").(*fiber.Group)
	routes.PublicRouting(public)

	log.Panic(app.Listen(":1200"))
}
