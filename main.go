package main

import (
	"note-dbs/controllers/usercontroller"
	"note-dbs/models"

	"github.com/gofiber/fiber/v2"
)

func main() {
	models.ConnectDB()
	app := fiber.New()

	api := app.Group("/api/user")
	api.Get("/", usercontroller.Index)
	api.Get("/:id", usercontroller.Show)
	api.Post("/", usercontroller.Create)
	api.Put("/:id", usercontroller.Update)
	api.Delete("/:id", usercontroller.Delete)

	app.Listen(":8000")
}
