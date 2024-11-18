package main

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/salawatbro/raxmet/config"
	"github.com/salawatbro/raxmet/database"
	middlewares "github.com/salawatbro/raxmet/internal/middleware"
	"github.com/salawatbro/raxmet/internal/routes"
	"github.com/salawatbro/raxmet/pkg/clients/redis"
	"github.com/salawatbro/raxmet/pkg/logger"
	"go.uber.org/zap"
)

func main() {
	//load config
	err := config.SetupConfig()
	if err != nil {
		panic(err)
	}
	//logger
	logger.InitLogger(config.Cfg.App.Env)
	defer logger.CloseLogger()
	//database
	err = database.ConnectDatabase()
	if err != nil {
		panic(err)
	}
	//redis
	redis.SetUpRedisClient()
	//new app
	app := newApp()
	//setup middlewares
	middlewares.Setup(app)
	//setup routes
	routes.ApiRoutes(app)
	//start server
	err = app.Listen(fmt.Sprintf("0.0.0.0:%s", config.Cfg.App.Port))
	if err != nil {
		logger.Logger.Error("Failed to start server", zap.Error(err))
	}
}

func newApp() *fiber.App {
	app := fiber.New(fiber.Config{
		JSONDecoder:    json.Unmarshal,
		JSONEncoder:    json.Marshal,
		ReadBufferSize: 4096 * 4,
		BodyLimit:      config.Cfg.App.MaxBody * 1024 * 1024,
	})

	return app
}
