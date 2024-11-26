package config

import (
	"os"
	"strings"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/pkg/errors"
)

type Config struct {
	Env       string    `json:"env" env-default:"local"`
	Logger    Logger    `json:"logger"`
	Kafka     Kafka     `json:"kafka"`
	AppConfig AppConfig `json:"frame_config"`
}

type AppConfig struct {
	HTTPServer                    HTTPServer `json:"http_server"`
	GracefulShutdownTimeoutSecond int64      `json:"graceful_shutdown_timeout_second"`
}

type HTTPServer struct {
	Port              int `json:"port"`
	ReadTimeoutSecond int `json:"read_timeout_second"`
}

type Kafka struct {
	WriteTimeoutSec int    `json:"write_timeout_sec" env:"KAFKA_WRITE_TIMEOUT_SEC" env-default:"120"`
	DSN             string `json:"dsn" env:"KAFKA_DSN"`
	ConsumerGroup   string `json:"consumer_group_id" env:"KAFKA_CONSUMER_GROUP"`
	Topics          struct {
		TestTopic string `json:"test-topic"`
	} `json:"topics"`
	Sasl struct {
		Enabled  bool   `json:"enabled" env:"KAFKA_SASL_ENABLED"`
		User     string `json:"user" env:"KAFKA_SASL_USER"`
		Password string `json:"password" env:"KAFKA_SASL_PASSWORD"`
		Cert     string `json:"cert" env:"KAFKA_SASL_CERT"`
	} `json:"sasl"`
}
type Logger struct {
	Level string `json:"level" env:"LOGGER_LEVEL"`
}

func CreateConfig() (*Config, error) {
	configPath := os.Getenv("CONFIG_PATH")
	if strings.TrimSpace(configPath) == "" {
		return nil, errors.New("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, errors.Errorf("config file does not exist - `%s`", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return nil, errors.Wrap(err, "cannot read config")
	}

	return &cfg, nil
}
