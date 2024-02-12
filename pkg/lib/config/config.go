package config

import (
	"github.com/caarlos0/env/v6"
	"golang.org/x/xerrors"
)

type Config struct {
}

func New() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, xerrors.Errorf("failed to parse cfg: %w", err)
	}
	return cfg, nil
}
