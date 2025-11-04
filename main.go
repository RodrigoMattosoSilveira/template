package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

const (
	ADMIN         = "admin"
)
func main() {
	// Create a new engine
	engine := html.New("./views", ".html")
	engine.AddFunc("isAdmin", func(user string) bool {
		return user == ADMIN
	})

	// Or from an embedded system
	// See github.com/gofiber/embed for examples
	// engine := html.NewFileSystem(http.Dir("./views"), ".html")

	// Pass the engine to the Views
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		// Render index
		return c.Render("index", fiber.Map{
			"Title": "Render index!",
		})
	})

	app.Get("/layout", func(c *fiber.Ctx) error {
		// Render index within layouts/main
		return c.Render("index", fiber.Map{
			"Title": "Render index within layouts/main!",
		}, "layouts/main")
	})

	app.Get("/layouts-nested", func(c *fiber.Ctx) error {
		// Render index within layouts/nested/main within layouts/nested/base
		return c.Render("index", fiber.Map{
			"Title": "Render index within layouts/nested/main within layouts/nested/base!",
		}, "layouts/nested/main", "layouts/nested/base")
	})

	log.Fatal(app.Listen(":3001"))
}
