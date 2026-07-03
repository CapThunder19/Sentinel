package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	Port         string
	Environment  string
	DatabaseURL  string
	RedisURL     string
	KafkaBrokers string
	JWTSecret    string
}

func Load() *Config {
	_ = godotenv.Load()
	viper.AutomaticEnv()

	cfg := &Config{
		Port:         viper.GetString("PORT"),
		Environment:  viper.GetString("APP_ENV"),
		DatabaseURL:  viper.GetString("DATABASE_URL"),
		RedisURL:     viper.GetString("REDIS_URL"),
		KafkaBrokers: viper.GetString("KAFKA_BROKERS"),
		JWTSecret:    viper.GetString("JWT_SECRET"),
	}

	if cfg.Port == "" {
		log.Fatal("PORT is required")
	}

	if cfg.JWTSecret == "" {
		log.Fatal("JWT_SECRET is required")
	}

	return cfg
}
