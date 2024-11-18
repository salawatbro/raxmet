package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/salawatbro/raxmet/internal/handlers"
	"github.com/salawatbro/raxmet/internal/repository"
	"github.com/salawatbro/raxmet/internal/services"
)

func ApiRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Not Found",
		})
	})
	//repositories
	userRepo := repository.NewUserRepository()
	//services
	authService := services.NewAuthService(userRepo)
	//handlers
	authHandler := handlers.NewAuthHandler(authService)

	api := app.Group("/api")
	api.Post("/login", authHandler.Login)
	api.Post("/register", authHandler.Register)
	//jwt
}
