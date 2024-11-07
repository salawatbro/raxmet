package config

import "github.com/spf13/viper"

type Config struct {
	App      AppConfig
	Server   ServerConfig
	Database DatabaseConfig
	Minio    MinioConfig
	Redis    RedisConfig
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
	Host string
	Port int
}

type ServerConfig struct {
	Port string
}

type AppConfig struct {
	Env     string
	Name    string
	Version string
	Debug   bool
}

func SetupConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
