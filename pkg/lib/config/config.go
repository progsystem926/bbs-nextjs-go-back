package config

import (
	"github.com/caarlos0/env/v6"
	"golang.org/x/xerrors"
)

type Config struct {
	Env        string `env:"API_ENV" envDefault:"dev"`
	DBHost     string `env:"API_DB_HOST" envDefault:"172.30.0.3"`
	DBName     string `env:"API_DB_NAME" envDefault:"golang"`
	DBUser     string `env:"API_DB_USER" envDefault:"root"`
	DBPass     string `env:"API_DB_PASS" envDefault:"pass"`
	DBPort     string `env:"API_DB_PORT" envDefault:"3306"`
	EncryptKey string `env:"API_ENCRYPT_KEY" envDefault:"passwordpassword"`
	SentryDsn  string `env:"API_SENTRY_DSN" envDefault:"https://xxxxxxxx"`
}

func New() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, xerrors.Errorf("failed to parse cfg: %w", err)
	}
	return cfg, nil
}
