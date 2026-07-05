package config

import (
	"log"

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
	viper.AutomaticEnv()

	return &Config{
		Port:         viper.GetString("PORT"),
		Environment:  viper.GetString("APP_ENV"),
		DatabaseURL:  viper.GetString("DATABASE_URL"),
		RedisURL:     viper.GetString("REDIS_URL"),
		KafkaBrokers: viper.GetString("KAFKA_BROKERS"),
		JWTSecret:    viper.GetString("JWT_SECRET"),
	}
}

func (c *Config) Validate() {
	if c.Port == "" {
		log.Fatal("PORT is required")
	}

	if c.DatabaseURL == "" {
		log.Fatal("DATABASE_URL is required")
	}

	if c.JWTSecret == "" {
		log.Fatal("JWT_SECRET is required")
	}
}
