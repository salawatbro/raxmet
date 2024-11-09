package routes

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/salawatbro/raxmet/config"
)

func ApiRoutes(app *fiber.App, cfg *config.Config) {
	api := app.Group("/api")
	api.Get("/v2", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	//jwt

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(cfg.JWT.Secret)},
	}))
	
	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}
