package configs

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/lmtani/learning-rate-limiter/internal/entity"
)

// Config struct to hold the configuration values
type Config struct {
	RedisAddr     string          `envconfig:"REDIS_ADDR"`
	RedisPassword string          `envconfig:"REDIS_PASSWORD"`
	WebServerPort string          `envconfig:"WEB_SERVER_PORT"`
	RateLimit     int             `envconfig:"RATE_LIMIT"`
	Expire        int             `envconfig:"EXPIRE"`
	TokenToLimit  entity.TokenMap `envconfig:"TOKEN_TO_LIMIT"`
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
