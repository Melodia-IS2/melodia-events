package config

import (
	"strings"

	"github.com/Melodia-IS2/melodia-go-utils/pkg/env"
	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	MongoConfig MongoConfig
	KafkaURL    string
	KafkaTopics []string
	RedisConfig RedisConfig
}

type MongoConfig struct {
	Port     string
	User     string
	Host     string
	Password string
	Database string
}

type RedisConfig struct {
	Host     string
	Port     string
	Password string
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
		RedisConfig: RedisConfig{
			Host:     env.GetEnv("REDIS_HOST", ""),
			Port:     env.GetEnv("REDIS_PORT", ""),
			Password: env.GetEnv("REDIS_PASSWORD", ""),
		},
		KafkaURL:    env.GetEnv("KAFKA_URL", ""),
		KafkaTopics: strings.Split(env.GetEnv("KAFKA_TOPICS", "__consumer_offsets"), ","),
	}
}
