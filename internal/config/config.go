package config

import (
	"github.com/Melodia-IS2/melodia-go-utils/pkg/env"
	"github.com/joho/godotenv"
)

type Config struct {
	Port string

	DatabaseURL string

	KafkaURL string
}

func Load() *Config {
	godotenv.Load()
	return &Config{
		Port: env.GetEnv("PORT", "8082"),

		DatabaseURL: env.GetEnv("DATABASE_URL", ""),

		KafkaURL: env.GetEnv("KAFKA_URL", ""),
	}
}
