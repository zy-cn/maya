package routers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// func NewRouter() *fiber.App {
// 	app := fiber.New()
// 	app.Static("/", "./public")

// 	apiV1 := app.Group("/api/v1")
// 	{
// 		apiV1.Get("/", func(c *fiber.Ctx) error {
// 			fmt.Println("zy")
// 			return nil
// 		})
// 	}

// 	return app

// }

func SetupRoutes(app *fiber.App) {
	apiV1 := app.Group("/api/v1")
	{
		apiV1.Get("/", func(c *fiber.Ctx) error {
			fmt.Println("zy")
			return nil
		})
	}
}
