package configs

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/lmtani/learning-rate-limiter/internal/entity"
)

// Config struct to hold the configuration values
type Config struct {
	RedisAddr         string          `envconfig:"REDIS_ADDR"`
	RedisPassword     string          `envconfig:"REDIS_PASSWORD"`
	WebServerPort     string          `envconfig:"WEB_SERVER_PORT"`
	RequestsPerSecond int             `envconfig:"REQUESTS_PER_SECOND"`
	WindowSize        int             `envconfig:"WINDOW_SIZE"`
	ApiKeyLimits      entity.TokenMap `envconfig:"API_KEY_LIMITS"`
}

func LoadConfig() (*Config, error) {
	var config Config

	// Carrega o arquivo .env
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	// Processa as variáveis de ambiente e preenche a struct Config
	if err := envconfig.Process("", &config); err != nil {
		return nil, err
	}

	return &config, nil
}
