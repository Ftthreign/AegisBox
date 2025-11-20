package config

import (
	"errors"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseConfig DatabaseConfig
	SessionConfig  SessionConfig
	SecurityConfig SecurityConfig
	LimiterConfig  LimiterConfig
	ServerConfig   ServerConfig
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	cfg := &Config{
		DatabaseConfig: loadDatabaseConfig(),
		SessionConfig:  loadSessionConfig(),
		SecurityConfig: loadSecurityConfig(),
		LimiterConfig:  loadLimiterConfig(),
		ServerConfig:   loadServerConfig(),
	}

	if err := cfg.validate(); err != nil {
		return nil, err
	}

	return cfg, nil
}

func (c *Config) validate() error {
	if len(c.SessionConfig.SessionSecret) < 32 {
		return errors.New("session secret must be at least 32 characters long")
	}

	return nil
}