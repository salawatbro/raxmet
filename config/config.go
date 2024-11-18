package config

import (
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	App      AppConfig
	Database DatabaseConfig
	Minio    MinioConfig
	Redis    RedisConfig
	JWT      JWTConfig
}

type DatabaseConfig struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     int
}

type MinioConfig struct {
	Endpoint  string
	AccessKey string
	SecretKey string
	Bucket    string
}

type RedisConfig struct {
	Host     string
	Port     int
	User     string
	Password string
}

type AppConfig struct {
	Port        string
	Env         string
	Name        string
	Version     string
	Debug       bool
	MaxBody     int
	MaxWorkers  int
	MaxRequests int
}

type JWTConfig struct {
	Secret string
	Exp    time.Duration
}

var Cfg *Config

func SetupConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return err
	}
	Cfg = &config
	return nil
}
