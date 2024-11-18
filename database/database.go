package database

import (
	"fmt"
	"go.uber.org/zap"
	"time"

	"github.com/salawatbro/raxmet/config"
	zaplogger "github.com/salawatbro/raxmet/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func ConnectDatabase() error {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Tashkent",
		config.Cfg.Database.Host, config.Cfg.Database.Port, config.Cfg.Database.User, config.Cfg.Database.Password, config.Cfg.Database.Name,
	)

	var gormLogger logger.Interface
	if config.Cfg.App.Env == "production" {
		gormLogger = logger.Default.LogMode(logger.Error)
	} else {
		gormLogger = logger.Default.LogMode(logger.Info)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		zaplogger.Logger.Error("Failed to connect to database", zap.Error(err))
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		zaplogger.Logger.Error("Failed to retrieve database object from GORM", zap.Error(err))
		return fmt.Errorf("failed to get db from GORM: %w", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	DB = db
	zaplogger.Logger.Info("Database connected successfully", zap.String("env", config.Cfg.App.Env))
	return nil
}
