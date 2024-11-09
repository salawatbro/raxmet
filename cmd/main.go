package main

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/salawatbro/raxmet/config"
	"github.com/salawatbro/raxmet/database"
	"github.com/salawatbro/raxmet/pkg/logger"
	"go.uber.org/zap"
)

func main() {
	cfg, err := config.SetupConfig()
	if err != nil {
		panic(err)
	}
	logger.InitLogger(cfg.App.Env)
	defer logger.CloseLogger()
	err = database.ConnectDatabase(cfg)
	if err != nil {
		panic(err)
	}
	app := newApp(cfg)
	err = app.Listen(fmt.Sprintf("0.0.0.0:%s", cfg.App.Port))
	if err != nil {
		logger.Logger.Error("Failed to start server", zap.Error(err))
	}
}

func newApp(cfg *config.Config) *fiber.App {
	app := fiber.New(fiber.Config{
		JSONDecoder:    json.Unmarshal,
		JSONEncoder:    json.Marshal,
		ReadBufferSize: 4096 * 4,
		BodyLimit:      cfg.App.MaxBody * 1024 * 1024,
	})

	return app
}
