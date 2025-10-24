package config

import (
	"github.com/Melodia-IS2/melodia-go-utils/pkg/env"
	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	MongoConfig MongoConfig
	KafkaURL    string
}

type MongoConfig struct {
	Port     string
	User     string
	Host     string
	Password string
	Database string
}

func Load() *Config {
	godotenv.Load()
	return &Config{
		Port: env.GetEnv("PORT", "8082"),

		MongoConfig: MongoConfig{
			Port:     env.GetEnv("MONGO_PORT", ""),
			User:     env.GetEnv("MONGO_USER", ""),
			Host:     env.GetEnv("MONGO_HOST", ""),
			Password: env.GetEnv("MONGO_PASSWORD", ""),
			Database: env.GetEnv("MONGO_DATABASE", ""),
		},

		KafkaURL: env.GetEnv("KAFKA_URL", ""),
	}
}
