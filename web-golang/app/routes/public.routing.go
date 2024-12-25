package routes

import (
	"belajar/web-golang/app/controller"

	"github.com/gofiber/fiber/v2"
)

func PublicRouting(app *fiber.Group) {
	app.Get("/", controller.Index)
	app.Get("/tentang", controller.Tentang)
	app.Get("/cokelat", controller.Coklat)
	app.Get("/testimoni", controller.Testimoni)
	app.Get("/kontak", controller.Kontak)
}
