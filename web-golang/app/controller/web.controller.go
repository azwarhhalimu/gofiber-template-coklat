package controller

import "github.com/gofiber/fiber/v2"

func Index(app *fiber.Ctx) error {
	return app.Render("index", fiber.Map{}, "layouts/main.layout")
}
func Tentang(app *fiber.Ctx) error {
	return app.Render("tentang", fiber.Map{}, "layouts/main.layout")
}
func Testimoni(app *fiber.Ctx) error {
	return app.Render("testimoni", fiber.Map{}, "layouts/main.layout")
}
func Coklat(app *fiber.Ctx) error {
	return app.Render("cokelat", fiber.Map{}, "layouts/main.layout")
}
func Kontak(app *fiber.Ctx) error {
	return app.Render("kontak", fiber.Map{}, "layouts/main.layout")
}
