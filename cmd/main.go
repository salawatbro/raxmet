package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/salawatbro/raxmet/config"
	"github.com/salawatbro/raxmet/pkg/logger"
	"go.uber.org/zap"
)

func main() {
	configs, err := config.SetupConfig()
	if err != nil {
		panic(err)
	}
	logger.InitLogger(configs.App.Env)
	defer logger.CloseLogger()
	app := fiber.New()

	err = app.Listen(fmt.Sprintf("localhost:%s", configs.Server.Port))
	if err != nil {
		logger.Logger.Error("Failed to start server", zap.Error(err))
	}
}
